// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"runtime/internal/sys"
	"unsafe"
)

func (s *smallmap) mapaccess1_fast64(t *maptype, key uint64) unsafe.Pointer {
	for i := 0; i < s.count; i++ {
		if key == s.key(t, i) {
			return s.elemptr(t, i)
		}
	}
	return unsafe.Pointer(&zeroVal[0])
}

func mapaccess1_fast64(t *maptype, c *mapcore, key uint64) unsafe.Pointer {
	if raceenabled && c != nil {
		callerpc := getcallerpc()
		racereadpc(unsafe.Pointer(c), callerpc, funcPC(mapaccess1_fast64))
	}
	if c == nil || c.count == 0 {
		return unsafe.Pointer(&zeroVal[0])
	}
	if c.flags&hashWriting != 0 {
		throw("concurrent map read and map write")
	}

	if c.impl() == mapImplSmall {
		s := (*smallmap)(unsafe.Pointer(c))
		return s.mapaccess1_fast64(t, key)
	}

	h := (*hmap)(unsafe.Pointer(c))
	var b *bmap
	if h.B == 0 {
		// One-bucket table. No need to hash.
		b = (*bmap)(h.buckets)
	} else {
		hash := t.key.alg.hash(noescape(unsafe.Pointer(&key)), uintptr(h.hash0))
		m := bucketMask(h.B)
		b = (*bmap)(add(h.buckets, (hash&m)*uintptr(t.bucketsize)))
		if c := h.oldbuckets; c != nil {
			if !h.sameSizeGrow() {
				// There used to be half as many buckets; mask down one more power of two.
				m >>= 1
			}
			oldb := (*bmap)(add(c, (hash&m)*uintptr(t.bucketsize)))
			if !evacuated(oldb) {
				b = oldb
			}
		}
	}
	for ; b != nil; b = b.overflow(t) {
		for i, k := uintptr(0), b.keys(); i < bucketCnt; i, k = i+1, add(k, 8) {
			if *(*uint64)(k) == key && !isEmpty(b.tophash[i]) {
				return add(unsafe.Pointer(b), dataOffset+bucketCnt*8+i*uintptr(t.valuesize))
			}
		}
	}
	return unsafe.Pointer(&zeroVal[0])
}

func (s *smallmap) mapaccess2_fast64(t *maptype, key uint64) (unsafe.Pointer, bool) {
	for i := 0; i < s.count; i++ {
		if key == s.key(t, i) {
			return s.elemptr(t, i), true
		}
	}
	return unsafe.Pointer(&zeroVal[0]), false
}

func mapaccess2_fast64(t *maptype, c *mapcore, key uint64) (unsafe.Pointer, bool) {
	if raceenabled && c != nil {
		callerpc := getcallerpc()
		racereadpc(unsafe.Pointer(c), callerpc, funcPC(mapaccess2_fast64))
	}
	if c == nil || c.count == 0 {
		return unsafe.Pointer(&zeroVal[0]), false
	}
	if c.flags&hashWriting != 0 {
		throw("concurrent map read and map write")
	}

	if c.impl() == mapImplSmall {
		s := (*smallmap)(unsafe.Pointer(c))
		return s.mapaccess2_fast64(t, key)
	}

	h := (*hmap)(unsafe.Pointer(c))
	var b *bmap
	if h.B == 0 {
		// One-bucket table. No need to hash.
		b = (*bmap)(h.buckets)
	} else {
		hash := t.key.alg.hash(noescape(unsafe.Pointer(&key)), uintptr(h.hash0))
		m := bucketMask(h.B)
		b = (*bmap)(add(h.buckets, (hash&m)*uintptr(t.bucketsize)))
		if c := h.oldbuckets; c != nil {
			if !h.sameSizeGrow() {
				// There used to be half as many buckets; mask down one more power of two.
				m >>= 1
			}
			oldb := (*bmap)(add(c, (hash&m)*uintptr(t.bucketsize)))
			if !evacuated(oldb) {
				b = oldb
			}
		}
	}
	for ; b != nil; b = b.overflow(t) {
		for i, k := uintptr(0), b.keys(); i < bucketCnt; i, k = i+1, add(k, 8) {
			if *(*uint64)(k) == key && !isEmpty(b.tophash[i]) {
				return add(unsafe.Pointer(b), dataOffset+bucketCnt*8+i*uintptr(t.valuesize)), true
			}
		}
	}
	return unsafe.Pointer(&zeroVal[0]), false
}

func (s *smallmap) promote64(t *maptype) {
	// make a shallow copy of s to work with
	s2 := *s
	// make a new map and (shallow) copy it over top of s
	*(*hmap)(unsafe.Pointer(s)) = *(makemap(t, s2.count, nil))
	// populate the new map with s2's contents
	for i := 0; i < s2.count; i++ {
		k := s2.key(t, i)
		e := s2.elemptr(t, i)
		// TODO: find some way to call directly into hmap mapassign
		// rather than having to go through the top-level dispatch mapassign_fast64?
		p := mapassign_fast64(t, (*mapcore)(unsafe.Pointer(s)), k)
		typedmemmove(t.elem, p, e)
	}
}

func (s *smallmap) mapassign_fast64(t *maptype, key uint64) unsafe.Pointer {
	if s.keys != nil && s.count == int(s.sz) {
		// full! promote.
		// (we could try to find out whether this assignment will overwrite
		// or insert, and not promote in the former case, but this is just a prototype.)
		s.promote64(t)
		// try again
		return mapassign_fast64(t, (*mapcore)(unsafe.Pointer(s)), key)
	}

	// there is room to add a new item if we need to
	s.flags ^= hashWriting
	var elem unsafe.Pointer
	if s.keys == nil {
		s.sz = 8
		s.keys = mallocgc(uintptr(s.sz)*uintptr(t.keysize), t.key, t.key.ptrdata != 0)
		s.elems = mallocgc(uintptr(s.sz)*uintptr(t.valuesize), t.elem, t.elem.ptrdata != 0)
	} else {
		for i := 0; i < s.count; i++ {
			if key == s.key(t, i) {
				elem = s.elemptr(t, i)
				break
			}
		}
	}
	if elem == nil {
		s.setkey(t, s.count, key)
		elem = s.elemptr(t, s.count)
		s.count++
	}
	if s.flags&hashWriting == 0 {
		throw("concurrent map writes")
	}
	s.flags &^= hashWriting
	return elem
}

func mapassign_fast64(t *maptype, c *mapcore, key uint64) unsafe.Pointer {
	if c == nil {
		panic(plainError("assignment to entry in nil map"))
	}
	if raceenabled {
		callerpc := getcallerpc()
		racewritepc(unsafe.Pointer(c), callerpc, funcPC(mapassign_fast64))
	}
	if c.flags&hashWriting != 0 {
		throw("concurrent map writes")
	}

	if c.impl() == mapImplSmall {
		s := (*smallmap)(unsafe.Pointer(c))
		return s.mapassign_fast64(t, key)
	}

	h := (*hmap)(unsafe.Pointer(c))
	hash := t.key.alg.hash(noescape(unsafe.Pointer(&key)), uintptr(h.hash0))

	// Set hashWriting after calling alg.hash for consistency with mapassign.
	h.flags ^= hashWriting

	if h.buckets == nil {
		h.buckets = newobject(t.bucket) // newarray(t.bucket, 1)
	}

again:
	bucket := hash & bucketMask(h.B)
	if h.growing() {
		growWork_fast64(t, h, bucket)
	}
	b := (*bmap)(unsafe.Pointer(uintptr(h.buckets) + bucket*uintptr(t.bucketsize)))

	var insertb *bmap
	var inserti uintptr
	var insertk unsafe.Pointer

bucketloop:
	for {
		for i := uintptr(0); i < bucketCnt; i++ {
			if isEmpty(b.tophash[i]) {
				if insertb == nil {
					insertb = b
					inserti = i
				}
				if b.tophash[i] == emptyRest {
					break bucketloop
				}
				continue
			}
			k := *((*uint64)(add(unsafe.Pointer(b), dataOffset+i*8)))
			if k != key {
				continue
			}
			insertb = b
			inserti = i
			goto done
		}
		ovf := b.overflow(t)
		if ovf == nil {
			break
		}
		b = ovf
	}

	// Did not find mapping for key. Allocate new cell & add entry.

	// If we hit the max load factor or we have too many overflow buckets,
	// and we're not already in the middle of growing, start growing.
	if !h.growing() && (overLoadFactor(h.count+1, h.B) || tooManyOverflowBuckets(h.noverflow, h.B)) {
		hashGrow(t, h)
		goto again // Growing the table invalidates everything, so try again
	}

	if insertb == nil {
		// all current buckets are full, allocate a new one.
		insertb = h.newoverflow(t, b)
		inserti = 0 // not necessary, but avoids needlessly spilling inserti
	}
	insertb.tophash[inserti&(bucketCnt-1)] = tophash(hash) // mask inserti to avoid bounds checks

	insertk = add(unsafe.Pointer(insertb), dataOffset+inserti*8)
	// store new key at insert position
	*(*uint64)(insertk) = key

	h.count++

done:
	val := add(unsafe.Pointer(insertb), dataOffset+bucketCnt*8+inserti*uintptr(t.valuesize))
	if h.flags&hashWriting == 0 {
		throw("concurrent map writes")
	}
	h.flags &^= hashWriting
	return val
}

func mapassign_fast64ptr(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer {
	if h == nil {
		panic(plainError("assignment to entry in nil map"))
	}
	if raceenabled {
		callerpc := getcallerpc()
		racewritepc(unsafe.Pointer(h), callerpc, funcPC(mapassign_fast64))
	}
	if h.flags&hashWriting != 0 {
		throw("concurrent map writes")
	}
	hash := t.key.alg.hash(noescape(unsafe.Pointer(&key)), uintptr(h.hash0))

	// Set hashWriting after calling alg.hash for consistency with mapassign.
	h.flags ^= hashWriting

	if h.buckets == nil {
		h.buckets = newobject(t.bucket) // newarray(t.bucket, 1)
	}

again:
	bucket := hash & bucketMask(h.B)
	if h.growing() {
		growWork_fast64(t, h, bucket)
	}
	b := (*bmap)(unsafe.Pointer(uintptr(h.buckets) + bucket*uintptr(t.bucketsize)))

	var insertb *bmap
	var inserti uintptr
	var insertk unsafe.Pointer

bucketloop:
	for {
		for i := uintptr(0); i < bucketCnt; i++ {
			if isEmpty(b.tophash[i]) {
				if insertb == nil {
					insertb = b
					inserti = i
				}
				if b.tophash[i] == emptyRest {
					break bucketloop
				}
				continue
			}
			k := *((*unsafe.Pointer)(add(unsafe.Pointer(b), dataOffset+i*8)))
			if k != key {
				continue
			}
			insertb = b
			inserti = i
			goto done
		}
		ovf := b.overflow(t)
		if ovf == nil {
			break
		}
		b = ovf
	}

	// Did not find mapping for key. Allocate new cell & add entry.

	// If we hit the max load factor or we have too many overflow buckets,
	// and we're not already in the middle of growing, start growing.
	if !h.growing() && (overLoadFactor(h.count+1, h.B) || tooManyOverflowBuckets(h.noverflow, h.B)) {
		hashGrow(t, h)
		goto again // Growing the table invalidates everything, so try again
	}

	if insertb == nil {
		// all current buckets are full, allocate a new one.
		insertb = h.newoverflow(t, b)
		inserti = 0 // not necessary, but avoids needlessly spilling inserti
	}
	insertb.tophash[inserti&(bucketCnt-1)] = tophash(hash) // mask inserti to avoid bounds checks

	insertk = add(unsafe.Pointer(insertb), dataOffset+inserti*8)
	// store new key at insert position
	*(*unsafe.Pointer)(insertk) = key

	h.count++

done:
	val := add(unsafe.Pointer(insertb), dataOffset+bucketCnt*8+inserti*uintptr(t.valuesize))
	if h.flags&hashWriting == 0 {
		throw("concurrent map writes")
	}
	h.flags &^= hashWriting
	return val
}

func (s *smallmap) mapdelete_fast64(t *maptype, key uint64) {
	s.flags ^= hashWriting
	for i := 0; i < s.count; i++ {
		if key == s.key(t, i) {
			endelem := s.elemptr(t, s.count-1)
			if i != s.count-1 {
				// copy final key/elem to this slot
				s.setkey(t, i, s.key(t, s.count-1))
				typedmemmove(t.elem, s.elemptr(t, i), endelem)
			}
			// no need to clear key
			// no need to clear value unless it has pointers
			if t.elem.ptrdata != 0 {
				memclrHasPointers(endelem, t.elem.size)
			}
			s.count--
			break
		}
	}
	if s.flags&hashWriting == 0 {
		throw("concurrent map writes")
	}
	s.flags &^= hashWriting
}

func mapdelete_fast64(t *maptype, c *mapcore, key uint64) {
	if raceenabled && c != nil {
		callerpc := getcallerpc()
		racewritepc(unsafe.Pointer(c), callerpc, funcPC(mapdelete_fast64))
	}
	if c == nil || c.count == 0 {
		return
	}
	if c.flags&hashWriting != 0 {
		throw("concurrent map writes")
	}

	if c.impl() == mapImplSmall {
		s := (*smallmap)(unsafe.Pointer(c))
		s.mapdelete_fast64(t, key)
		return
	}

	h := (*hmap)(unsafe.Pointer(c))

	hash := t.key.alg.hash(noescape(unsafe.Pointer(&key)), uintptr(h.hash0))

	// Set hashWriting after calling alg.hash for consistency with mapdelete
	h.flags ^= hashWriting

	bucket := hash & bucketMask(h.B)
	if h.growing() {
		growWork_fast64(t, h, bucket)
	}
	b := (*bmap)(add(h.buckets, bucket*uintptr(t.bucketsize)))
	bOrig := b
search:
	for ; b != nil; b = b.overflow(t) {
		for i, k := uintptr(0), b.keys(); i < bucketCnt; i, k = i+1, add(k, 8) {
			if key != *(*uint64)(k) || isEmpty(b.tophash[i]) {
				continue
			}
			// Only clear key if there are pointers in it.
			if t.key.ptrdata != 0 {
				memclrHasPointers(k, t.key.size)
			}
			v := add(unsafe.Pointer(b), dataOffset+bucketCnt*8+i*uintptr(t.valuesize))
			if t.elem.ptrdata != 0 {
				memclrHasPointers(v, t.elem.size)
			} else {
				memclrNoHeapPointers(v, t.elem.size)
			}
			b.tophash[i] = emptyOne
			// If the bucket now ends in a bunch of emptyOne states,
			// change those to emptyRest states.
			if i == bucketCnt-1 {
				if b.overflow(t) != nil && b.overflow(t).tophash[0] != emptyRest {
					goto notLast
				}
			} else {
				if b.tophash[i+1] != emptyRest {
					goto notLast
				}
			}
			for {
				b.tophash[i] = emptyRest
				if i == 0 {
					if b == bOrig {
						break // beginning of initial bucket, we're done.
					}
					// Find previous bucket, continue at its last entry.
					c := b
					for b = bOrig; b.overflow(t) != c; b = b.overflow(t) {
					}
					i = bucketCnt - 1
				} else {
					i--
				}
				if b.tophash[i] != emptyOne {
					break
				}
			}
		notLast:
			h.count--
			break search
		}
	}

	if h.flags&hashWriting == 0 {
		throw("concurrent map writes")
	}
	h.flags &^= hashWriting
}

func growWork_fast64(t *maptype, h *hmap, bucket uintptr) {
	// make sure we evacuate the oldbucket corresponding
	// to the bucket we're about to use
	evacuate_fast64(t, h, bucket&h.oldbucketmask())

	// evacuate one more oldbucket to make progress on growing
	if h.growing() {
		evacuate_fast64(t, h, h.nevacuate)
	}
}

func evacuate_fast64(t *maptype, h *hmap, oldbucket uintptr) {
	b := (*bmap)(add(h.oldbuckets, oldbucket*uintptr(t.bucketsize)))
	newbit := h.noldbuckets()
	if !evacuated(b) {
		// TODO: reuse overflow buckets instead of using new ones, if there
		// is no iterator using the old buckets.  (If !oldIterator.)

		// xy contains the x and y (low and high) evacuation destinations.
		var xy [2]evacDst
		x := &xy[0]
		x.b = (*bmap)(add(h.buckets, oldbucket*uintptr(t.bucketsize)))
		x.k = add(unsafe.Pointer(x.b), dataOffset)
		x.v = add(x.k, bucketCnt*8)

		if !h.sameSizeGrow() {
			// Only calculate y pointers if we're growing bigger.
			// Otherwise GC can see bad pointers.
			y := &xy[1]
			y.b = (*bmap)(add(h.buckets, (oldbucket+newbit)*uintptr(t.bucketsize)))
			y.k = add(unsafe.Pointer(y.b), dataOffset)
			y.v = add(y.k, bucketCnt*8)
		}

		for ; b != nil; b = b.overflow(t) {
			k := add(unsafe.Pointer(b), dataOffset)
			v := add(k, bucketCnt*8)
			for i := 0; i < bucketCnt; i, k, v = i+1, add(k, 8), add(v, uintptr(t.valuesize)) {
				top := b.tophash[i]
				if isEmpty(top) {
					b.tophash[i] = evacuatedEmpty
					continue
				}
				if top < minTopHash {
					throw("bad map state")
				}
				var useY uint8
				if !h.sameSizeGrow() {
					// Compute hash to make our evacuation decision (whether we need
					// to send this key/value to bucket x or bucket y).
					hash := t.key.alg.hash(k, uintptr(h.hash0))
					if hash&newbit != 0 {
						useY = 1
					}
				}

				b.tophash[i] = evacuatedX + useY // evacuatedX + 1 == evacuatedY, enforced in makemap
				dst := &xy[useY]                 // evacuation destination

				if dst.i == bucketCnt {
					dst.b = h.newoverflow(t, dst.b)
					dst.i = 0
					dst.k = add(unsafe.Pointer(dst.b), dataOffset)
					dst.v = add(dst.k, bucketCnt*8)
				}
				dst.b.tophash[dst.i&(bucketCnt-1)] = top // mask dst.i as an optimization, to avoid a bounds check

				// Copy key.
				if t.key.ptrdata != 0 && writeBarrier.enabled {
					if sys.PtrSize == 8 {
						// Write with a write barrier.
						*(*unsafe.Pointer)(dst.k) = *(*unsafe.Pointer)(k)
					} else {
						// There are three ways to squeeze at least one 32 bit pointer into 64 bits.
						// Give up and call typedmemmove.
						typedmemmove(t.key, dst.k, k)
					}
				} else {
					*(*uint64)(dst.k) = *(*uint64)(k)
				}

				typedmemmove(t.elem, dst.v, v)
				dst.i++
				// These updates might push these pointers past the end of the
				// key or value arrays.  That's ok, as we have the overflow pointer
				// at the end of the bucket to protect against pointing past the
				// end of the bucket.
				dst.k = add(dst.k, 8)
				dst.v = add(dst.v, uintptr(t.valuesize))
			}
		}
		// Unlink the overflow buckets & clear key/value to help GC.
		if h.flags&oldIterator == 0 && t.bucket.ptrdata != 0 {
			b := add(h.oldbuckets, oldbucket*uintptr(t.bucketsize))
			// Preserve b.tophash because the evacuation
			// state is maintained there.
			ptr := add(b, dataOffset)
			n := uintptr(t.bucketsize) - dataOffset
			memclrHasPointers(ptr, n)
		}
	}

	if oldbucket == h.nevacuate {
		advanceEvacuationMark(h, t, newbit)
	}
}
