// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mime implements parts of the MIME spec.
package mime

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
)

var (
	mimeTypesValue      atomic.Value // of map[string]string
	mimeTypesLowerValue atomic.Value // of map[string]string
	mimemu              sync.Mutex   // serializes mime type additions
	once                sync.Once    // guards initMime
)

func mimeTypes() map[string]string      { return mimeTypesValue.Load().(map[string]string) }
func mimeTypesLower() map[string]string { return mimeTypesLowerValue.Load().(map[string]string) }

func clone(m map[string]string) map[string]string {
	m2 := make(map[string]string, len(m))
	for k, v := range m {
		m2[k] = v
	}
	return m2
}

func initMime() {
	baseMimeTypes := map[string]string{
		".css":  "text/css; charset=utf-8",
		".gif":  "image/gif",
		".htm":  "text/html; charset=utf-8",
		".html": "text/html; charset=utf-8",
		".jpg":  "image/jpeg",
		".js":   "application/x-javascript",
		".pdf":  "application/pdf",
		".png":  "image/png",
		".xml":  "text/xml; charset=utf-8",
	}
	for k := range baseMimeTypes {
		if strings.ToLower(k) != k {
			panic("keys in baseMimeTypes must be lowercase")
		}
	}

	baseMimeTypesLower := clone(baseMimeTypes)
	for ext, typ := range initMimePlatform() {
		typ, err := formatMimeType(typ)
		if err != nil {
			continue
		}
		baseMimeTypes[ext] = typ
		baseMimeTypesLower[strings.ToLower(ext)] = typ
	}

	mimeTypesValue.Store(baseMimeTypes)
	mimeTypesLowerValue.Store(baseMimeTypesLower)
}

// TypeByExtension returns the MIME type associated with the file extension ext.
// The extension ext should begin with a leading dot, as in ".html".
// When ext has no associated type, TypeByExtension returns "".
//
// Extensions are looked up first case-sensitively, then case-insensitively.
//
// The built-in table is small but on unix it is augmented by the local
// system's mime.types file(s) if available under one or more of these
// names:
//
//   /etc/mime.types
//   /etc/apache2/mime.types
//   /etc/apache/mime.types
//
// On Windows, MIME types are extracted from the registry.
//
// Text types have the charset parameter set to "utf-8" by default.
func TypeByExtension(ext string) string {
	once.Do(initMime)

	// Case-sensitive lookup.
	v := mimeTypes()[ext]
	if v != "" {
		return v
	}

	// Case-insensitive lookup.
	// Optimistically assume a short ASCII extension and be
	// allocation-free in that case.
	var buf [10]byte
	lower := buf[:0]
	const utf8RuneSelf = 0x80 // from utf8 package, but not importing it.
	for i := 0; i < len(ext); i++ {
		c := ext[i]
		if c >= utf8RuneSelf {
			// Slow path.
			return mimeTypesLower()[strings.ToLower(ext)]
		}
		if 'A' <= c && c <= 'Z' {
			lower = append(lower, c+('a'-'A'))
		} else {
			lower = append(lower, c)
		}
	}
	// The conversion from []byte to string doesn't allocate in
	// a map lookup.
	return mimeTypesLower()[string(lower)]
}

// AddExtensionType sets the MIME type associated with
// the extension ext to typ. The extension should begin with
// a leading dot, as in ".html".
func AddExtensionType(ext, typ string) error {
	if !strings.HasPrefix(ext, ".") {
		return fmt.Errorf(`mime: extension %q misses dot`, ext)
	}
	once.Do(initMime)
	mimemu.Lock()
	defer mimemu.Unlock()
	return setExtensionType(ext, typ)
}

func formatMimeType(mimeType string) (string, error) {
	_, param, err := ParseMediaType(mimeType)
	if err != nil {
		return "", err
	}
	if strings.HasPrefix(mimeType, "text/") && param["charset"] == "" {
		param["charset"] = "utf-8"
		mimeType = FormatMediaType(mimeType, param)
	}
	return mimeType, nil
}

func setExtensionType(extension, mimeType string) error {
	mimeType, err := formatMimeType(mimeType)
	if err != nil {
		return err
	}

	new := clone(mimeTypes())
	new[extension] = mimeType
	mimeTypesValue.Store(new)

	new = clone(mimeTypesLower())
	new[strings.ToLower(extension)] = mimeType
	mimeTypesLowerValue.Store(new)
	return nil
}
