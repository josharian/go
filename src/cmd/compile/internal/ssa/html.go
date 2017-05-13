// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"bytes"
	"cmd/internal/src"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

type HTMLWriter struct {
	Logger
	w   io.WriteCloser
	dot *dotWriter
}

func NewHTMLWriter(path string, logger Logger, funcname string) *HTMLWriter {
	out, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		logger.Fatalf(src.NoXPos, "%v", err)
	}
	html := HTMLWriter{w: out, Logger: logger}
	html.dot = newDotWriter()
	html.start(funcname)
	return &html
}

func (w *HTMLWriter) start(name string) {
	if w == nil {
		return
	}
	w.WriteString("<html>")
	// TODO: These numbers work well for fannkuch.
	// The columns are too big for simpler CFGs.
	// How do I pick a good size?
	// And it will need to be applied post-facto;
	// should we buffer the entire HTML so that
	// we can fix it up in html head,
	// or should we fix it with javascript?
	// If we fix it with javascript,
	// we can just let the user pick the size.
	// This seems better but the resulting reflow
	// seems to make Chrome lock up.
	tableWidth := "400"
	elemWidth := "300"
	if w.dot.err == nil {
		tableWidth = "800"
		elemWidth = "700"
	}
	w.WriteString(`<head>
<meta http-equiv="Content-Type" content="text/html;charset=UTF-8">
<style>

#helplink {
    margin-bottom: 15px;
    display: block;
    margin-top: -15px;
}

#help {
    display: none;
}

.stats {
	font-size: 60%;
}

table {
    border: 1px solid black;
    table-layout: fixed;
    width: ` + tableWidth + `px;
}

th, td {
    border: 1px solid black;
    overflow: hidden;
    width: ` + elemWidth + `px;
    vertical-align: top;
    padding: 5px;
}

li {
    list-style-type: none;
}

li.ssa-long-value {
    text-indent: -2em;  /* indent wrapped lines */
}

li.ssa-value-list {
    display: inline;
}

li.ssa-start-block {
    padding: 0;
    margin: 0;
}

li.ssa-end-block {
    padding: 0;
    margin: 0;
}

ul.ssa-print-func {
    padding-left: 0;
}

dl.ssa-gen {
    padding-left: 0;
}

dt.ssa-prog-src {
    padding: 0;
    margin: 0;
    float: left;
    width: 4em;
}

dd.ssa-prog {
    padding: 0;
    margin-right: 0;
    margin-left: 4em;
}

.dead-value {
    color: gray;
}

.dead-block {
    opacity: 0.5;
}

.depcycle {
    font-style: italic;
}

.highlight-yellow         { background-color: yellow; }
.highlight-aquamarine     { background-color: aquamarine; }
.highlight-coral          { background-color: coral; }
.highlight-lightpink      { background-color: lightpink; }
.highlight-lightsteelblue { background-color: lightsteelblue; }
.highlight-palegreen      { background-color: palegreen; }
.highlight-powderblue     { background-color: powderblue; }
.highlight-lightgray      { background-color: lightgray; }

span.outline-blue           { outline: blue solid 2px; }
span.outline-red            { outline: red solid 2px; }
span.outline-blueviolet     { outline: blueviolet solid 2px; }
span.outline-darkolivegreen { outline: darkolivegreen solid 2px; }
span.outline-fuchsia        { outline: fuchsia solid 2px; }
span.outline-sienna         { outline: sienna solid 2px; }
span.outline-gold           { outline: gold solid 2px; }

ellipse.outline-blue           { stroke: blue; stroke-width: 3; }
ellipse.outline-red            { stroke: red; stroke-width: 3; }
ellipse.outline-blueviolet     { stroke: blueviolet; stroke-width: 3; }
ellipse.outline-darkolivegreen { stroke: darkolivegreen; stroke-width: 3; }
ellipse.outline-fuchsia        { stroke: fuchsia; stroke-width: 3; }
ellipse.outline-sienna         { stroke: sienna; stroke-width: 3; }
ellipse.outline-gold           { stroke: gold; stroke-width: 3; }

</style>

<script type="text/javascript">
// ordered list of all available highlight colors
var highlights = [
    "highlight-aquamarine",
    "highlight-coral",
    "highlight-lightpink",
    "highlight-lightsteelblue",
    "highlight-palegreen",
    "highlight-lightgray",
    "highlight-yellow"
];

// state: which value is highlighted this color?
var highlighted = {};
for (var i = 0; i < highlights.length; i++) {
    highlighted[highlights[i]] = "";
}

// ordered list of all available outline colors
var outlines = [
    "outline-blue",
    "outline-red",
    "outline-blueviolet",
    "outline-darkolivegreen",
    "outline-fuchsia",
    "outline-sienna",
    "outline-gold"
];

// state: which value is outlined this color?
var outlined = {};
for (var i = 0; i < outlines.length; i++) {
    outlined[outlines[i]] = "";
}

window.onload = function() {
    var ssaElemClicked = function(elem, event, selections, selected) {
        event.stopPropagation()

        // TODO: pushState with updated state and read it on page load,
        // so that state can survive across reloads

        // find all values with the same name
        var c = elem.classList.item(0);
        var x = document.getElementsByClassName(c);

        // if selected, remove selections from all of them
        // otherwise, attempt to add

        var remove = "";
        for (var i = 0; i < selections.length; i++) {
            var color = selections[i];
            if (selected[color] == c) {
                remove = color;
                break;
            }
        }

        if (remove != "") {
            for (var i = 0; i < x.length; i++) {
                x[i].classList.remove(remove);
            }
            selected[remove] = "";
            return;
        }

        // we're adding a selection
        // find first available color
        var avail = "";
        for (var i = 0; i < selections.length; i++) {
            var color = selections[i];
            if (selected[color] == "") {
                avail = color;
                break;
            }
        }
        if (avail == "") {
            alert("out of selection colors; go add more");
            return;
        }

        // set that as the selection
        for (var i = 0; i < x.length; i++) {
            x[i].classList.add(avail);
        }
        selected[avail] = c;
    };

    var ssaValueClicked = function(event) {
        ssaElemClicked(this, event, highlights, highlighted);
    }

    var ssaBlockClicked = function(event) {
        ssaElemClicked(this, event, outlines, outlined);
    }

    var ssavalues = document.getElementsByClassName("ssa-value");
    for (var i = 0; i < ssavalues.length; i++) {
        ssavalues[i].addEventListener('click', ssaValueClicked);
    }

    var ssalongvalues = document.getElementsByClassName("ssa-long-value");
    for (var i = 0; i < ssalongvalues.length; i++) {
        // don't attach listeners to li nodes, just the spans they contain
        if (ssalongvalues[i].nodeName == "SPAN") {
            ssalongvalues[i].addEventListener('click', ssaValueClicked);
        }
    }

    var ssablocks = document.getElementsByClassName("ssa-block");
    for (var i = 0; i < ssablocks.length; i++) {
        ssablocks[i].addEventListener('click', ssaBlockClicked);
    }

    // find all svg block nodes, add their block classes
    var nodes = document.querySelectorAll('*[id^="graph_node_"]');
    for (var i = 0; i < nodes.length; i++) {
    	var node = nodes[i];
    	var name = node.id.toString();
    	var block = name.substring(name.lastIndexOf("_")+1);
    	node.classList.remove("node");
    	node.classList.add(block);
        node.addEventListener('click', ssaBlockClicked);
        var ellipse = node.getElementsByTagName('ellipse')[0];
        ellipse.classList.add(block);
    }

    document.onkeypress = function(e) {
    	console.log(e.keyCode);
    	return; // TODO: decide what to do here...see comments about table width above
        switch (e.keyCode) {
        case 'w'.charCodeAt():
        	// Make columns wider by applying a new "wide columns" class.
        	var tagnames = ["table", "th", "td"];
        	for (var j = 0; j < tagnames.length; i++) {
        		console.log("tag", tagnames[j])
        		var x = document.getElementsByTagName(tagnames[j]);
		        for (var i = 0; i < x.length; i++) {
		        	console.log("add width3 to", x[i])
		            x[i].classList.add("width3");
		        }
        	}
        case 's'.charCodeAt():
        	// TODO: make skinnier
        }
    };
};

function toggle_visibility(id) {
   var e = document.getElementById(id);
   if(e.style.display == 'block')
      e.style.display = 'none';
   else
      e.style.display = 'block';
}
</script>

</head>`)
	w.WriteString("<body>")
	w.WriteString("<h1>")
	w.WriteString(html.EscapeString(name))
	w.WriteString("</h1>")
	w.WriteString(`
<a href="#" onclick="toggle_visibility('help');" id="helplink">help</a>
<div id="help">

<p>
Click on a value or block to toggle highlighting of that value/block
and its uses.  (Values and blocks are highlighted by ID, and IDs of
dead items may be reused, so not all highlights necessarily correspond
to the clicked item.)
</p>

<p>
Faded out values and blocks are dead code that has not been eliminated.
</p>

<p>
Values printed in italics have a dependency cycle.
</p>

<p>
Press 'w' to make the columns wider, 's' to make them skinnier.
</pr>

</div>
`)
	w.WriteString("<table>")
	w.WriteString("<tr>")
}

func (w *HTMLWriter) Close() {
	if w == nil {
		return
	}
	io.WriteString(w.w, "</tr>")
	io.WriteString(w.w, "</table>")
	io.WriteString(w.w, "</body>")
	io.WriteString(w.w, "</html>")
	// if w.dot.err != nil {
	// 	// TODO: Put this somewhere visible in the HTML instead of panicking
	// 	panic(w.dot.err)
	// }
	w.w.Close()
}

// WriteFunc writes f in a column headed by title.
func (w *HTMLWriter) WriteFunc(title string, f *Func) {
	if w == nil {
		return // avoid generating HTML just to discard it
	}
	w.WriteColumn(title, f.HTML(w.dot))
}

// WriteColumn writes raw HTML in a column headed by title.
// It is intended for pre- and post-compilation log output.
func (w *HTMLWriter) WriteColumn(title string, html string) {
	if w == nil {
		return
	}
	w.WriteString("<td>")
	w.WriteString("<h2>" + title + "</h2>")
	w.WriteString(html)
	w.WriteString("</td>")
}

func (w *HTMLWriter) Printf(msg string, v ...interface{}) {
	if _, err := fmt.Fprintf(w.w, msg, v...); err != nil {
		w.Fatalf(src.NoXPos, "%v", err)
	}
}

func (w *HTMLWriter) WriteString(s string) {
	if _, err := io.WriteString(w.w, s); err != nil {
		w.Fatalf(src.NoXPos, "%v", err)
	}
}

func (v *Value) HTML() string {
	// TODO: Using the value ID as the class ignores the fact
	// that value IDs get recycled and that some values
	// are transmuted into other values.
	s := v.String()
	return fmt.Sprintf("<span class=\"%s ssa-value\">%s</span>", s, s)
}

func (v *Value) LongHTML() string {
	// TODO: Any intra-value formatting?
	// I'm wary of adding too much visual noise,
	// but a little bit might be valuable.
	// We already have visual noise in the form of punctuation
	// maybe we could replace some of that with formatting.
	s := fmt.Sprintf("<span class=\"%s ssa-long-value\">", v.String())
	s += fmt.Sprintf("%s = %s", v.HTML(), v.Op.String())
	s += " &lt;" + html.EscapeString(v.Type.String()) + "&gt;"
	s += html.EscapeString(v.auxString())
	for _, a := range v.Args {
		s += fmt.Sprintf(" %s", a.HTML())
	}
	r := v.Block.Func.RegAlloc
	if int(v.ID) < len(r) && r[v.ID] != nil {
		s += " : " + html.EscapeString(r[v.ID].Name())
	}
	s += "</span>"
	return s
}

func (b *Block) HTML() string {
	// TODO: Using the value ID as the class ignores the fact
	// that value IDs get recycled and that some values
	// are transmuted into other values.
	s := html.EscapeString(b.String())
	return fmt.Sprintf("<span class=\"%s ssa-block\">%s</span>", s, s)
}

func (b *Block) LongHTML() string {
	// TODO: improve this for HTML?
	s := fmt.Sprintf("<span class=\"%s ssa-block\">%s</span>", html.EscapeString(b.String()), html.EscapeString(b.Kind.String()))
	if b.Aux != nil {
		s += html.EscapeString(fmt.Sprintf(" {%v}", b.Aux))
	}
	if b.Control != nil {
		s += fmt.Sprintf(" %s", b.Control.HTML())
	}
	if len(b.Succs) > 0 {
		s += " &#8594;" // right arrow
		for _, e := range b.Succs {
			c := e.b
			s += " " + c.HTML()
		}
	}
	switch b.Likely {
	case BranchUnlikely:
		s += " (unlikely)"
	case BranchLikely:
		s += " (likely)"
	}
	return s
}

func (f *Func) HTML(dot *dotWriter) string {
	buf := new(bytes.Buffer)
	dot.writeFuncSVG(buf, f)
	fmt.Fprint(buf, "<code>")
	p := htmlFuncPrinter{w: buf}
	fprintFunc(p, f)

	// fprintFunc(&buf, f) // TODO: HTML, not text, <br /> for line breaks, etc.
	fmt.Fprint(buf, "</code>")
	return buf.String()
}

func (d *dotWriter) writeFuncSVG(w io.Writer, f *Func) {
	if d.err != nil {
		return
	}
	buf := new(bytes.Buffer)
	cmd := exec.Command(d.path, "-Tsvg")
	pipe, err := cmd.StdinPipe()
	d.setErr(err)
	cmd.Stdout = buf
	d.setErr(cmd.Start())
	fmt.Fprintln(pipe, "digraph {")
	//fmt.Fprintln(pipe, "splines=ortho;")
	for i, b := range f.Blocks {
		layout := ""
		if f.laidout {
			layout = fmt.Sprintf(" (%d)", i)
		}
		fmt.Fprintf(pipe, `%v [label="%v%s\n%v",id="graph_node_%d_%v"];`, b, b, layout, b.Kind, d.ngraphs, b)
		fmt.Fprintln(pipe)
	}
	indexOf := make([]int, f.NumBlocks())
	for i, b := range f.Blocks {
		indexOf[b.ID] = i
	}
	layoutDrawn := make([]bool, f.NumBlocks())
	for _, b := range f.Blocks {
		for i, s := range b.Succs {
			style := "solid"
			if b.unlikelyIndex() == i {
				style = "dashed"
			}
			color := "black"
			if f.laidout && indexOf[s.b.ID] == indexOf[b.ID]+1 {
				color = "green"
				layoutDrawn[s.b.ID] = true
			}
			fmt.Fprintf(pipe, `%v -> %v [label=" %d ",style="%s",color="%s"];`, b, s.b, i, style, color)
			fmt.Fprintln(pipe)
		}
	}
	if f.laidout {
		fmt.Fprintln(pipe, "edge[constraint=false];")
		for i := 1; i < len(f.Blocks); i++ {
			if layoutDrawn[f.Blocks[i].ID] {
				continue
			}
			fmt.Fprintf(pipe, `%s -> %s [color="green",style="dotted"];`, f.Blocks[i-1], f.Blocks[i])
			fmt.Fprintln(pipe)
		}
	}
	fmt.Fprintln(pipe, "}")
	pipe.Close()
	d.setErr(cmd.Wait())

	// Apparently there's no way to give a reliable target width to dot?
	// And no way to supply an HTML class for the svg element either?
	// For now, use an awful hack--edit the html as it passes through
	// our fingers, finding '<svg width="..." height="..." [everything else]'
	// and replacing it with '<svg width="100%" [everything else]'.
	d.copyAfter(w, buf, `<svg `)
	d.copyAfter(w, buf, `width="`)
	io.WriteString(w, `100%"`)
	d.copyAfter(ioutil.Discard, buf, `"`)
	d.copyAfter(ioutil.Discard, buf, `height="`)
	d.copyAfter(ioutil.Discard, buf, `"`)
	if d.err != nil {
		return
	}
	io.Copy(w, buf)
	d.ngraphs++ // used to give each node a unique id
}

func (b *Block) unlikelyIndex() int {
	switch b.Likely {
	case BranchLikely:
		return 1
	case BranchUnlikely:
		return 0
	}
	return -1
}

func (d *dotWriter) copyAfter(w io.Writer, buf *bytes.Buffer, sep string) {
	if d.err != nil {
		return
	}
	i := bytes.Index(buf.Bytes(), []byte(sep))
	if i == -1 {
		d.setErr(fmt.Errorf("couldn't find dot sep %q", sep))
		return
	}
	io.CopyN(w, buf, int64(i+len(sep)))
}

type htmlFuncPrinter struct {
	w io.Writer
}

func (p htmlFuncPrinter) header(f *Func) {}

func (p htmlFuncPrinter) startBlock(b *Block, reachable bool) {
	// TODO: Make blocks collapsable?
	var dead string
	if !reachable {
		dead = "dead-block"
	}
	fmt.Fprintf(p.w, "<ul class=\"%s ssa-print-func %s\">", b, dead)
	fmt.Fprintf(p.w, "<li class=\"ssa-start-block\">%s:", b.HTML())
	if len(b.Preds) > 0 {
		io.WriteString(p.w, " &#8592;") // left arrow
		for _, e := range b.Preds {
			pred := e.b
			fmt.Fprintf(p.w, " %s", pred.HTML())
		}
	}
	io.WriteString(p.w, "</li>")
	if len(b.Values) > 0 { // start list of values
		io.WriteString(p.w, "<li class=\"ssa-value-list\">")
		io.WriteString(p.w, "<ul>")
	}
}

func (p htmlFuncPrinter) endBlock(b *Block) {
	if len(b.Values) > 0 { // end list of values
		io.WriteString(p.w, "</ul>")
		io.WriteString(p.w, "</li>")
	}
	io.WriteString(p.w, "<li class=\"ssa-end-block\">")
	fmt.Fprint(p.w, b.LongHTML())
	io.WriteString(p.w, "</li>")
	io.WriteString(p.w, "</ul>")
	// io.WriteString(p.w, "</span>")
}

func (p htmlFuncPrinter) value(v *Value, live bool) {
	var dead string
	if !live {
		dead = "dead-value"
	}
	fmt.Fprintf(p.w, "<li class=\"ssa-long-value %s\">", dead)
	fmt.Fprint(p.w, v.LongHTML())
	io.WriteString(p.w, "</li>")
}

func (p htmlFuncPrinter) startDepCycle() {
	fmt.Fprintln(p.w, "<span class=\"depcycle\">")
}

func (p htmlFuncPrinter) endDepCycle() {
	fmt.Fprintln(p.w, "</span>")
}

func (p htmlFuncPrinter) named(n LocalSlot, vals []*Value) {
	fmt.Fprintf(p.w, "<li>name %s: ", n.Name())
	for _, val := range vals {
		fmt.Fprintf(p.w, "%s ", val.HTML())
	}
	fmt.Fprintf(p.w, "</li>")
}

type dotWriter struct {
	path    string
	err     error
	ngraphs int
}

func newDotWriter() *dotWriter {
	// if os.Getenv("GOSSACFG") == "" {
	// 	return &dotWriter{err: errors.New("enable visual CFG by setting GOSSACFG=1")}
	// }
	path, err := exec.LookPath("dot")
	return &dotWriter{path: path, err: err}
}

func (d *dotWriter) setErr(err error) {
	if err == nil {
		return
	}
	if d.err == nil {
		d.err = err
	}
}
