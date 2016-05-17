package main

import (
	"cmd/internal/traceviewer"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

func writeTraceHTML(w io.Writer, data traceviewer.Data) error {
	// Write all the parts of the HMTL output, one at a time.
	// See the const block below.
	_, err := io.WriteString(w, traceHTML1)
	if err != nil {
		return err
	}
	frag, err := os.Open(filepath.Join(runtime.GOROOT(), "misc", "trace", "trace_viewer_lean_fragment.html"))
	if err != nil {
		return err
	}
	_, err = io.Copy(w, frag)
	if err != nil {
		return err
	}
	_, err = io.WriteString(w, traceHTML2)
	if err != nil {
		return err
	}
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		return err
	}
	_, err = io.WriteString(w, traceHTML3)
	if err != nil {
		return err
	}
	return nil
}

// Snippets used to generate a trace html file.
// Follows https://github.com/catapult-project/catapult/blob/master/tracing/docs/embedding-trace-viewer.md
const (
	traceHTML1 = `<html><head>`
	// Then trace viewer fragment loaded from file
	traceHTML2 = `
<script>
(function() {
  document.addEventListener('DOMContentLoaded', function() {
    var container = document.createElement('track-view-container');
    container.id = 'track_view_container';

    var viewer = document.createElement('tr-ui-timeline-view');
    viewer.track_view_container = container;
    viewer.appendChild(container);

    viewer.id = 'trace-viewer';
    viewer.globalMode = true;
    document.body.appendChild(viewer);
    var result = 
`
	// Then trace JSON
	traceHTML3 = `;
    var model = new tr.Model();
    var i = new tr.importer.Import(model);
    // TODO: The progress dialog is overkill here. Import synchronously instead.
    var p = i.importTracesWithProgressDialog([result]);
    p.then(
    	// success
    	function() {
    		viewer.model = model;
    		viewer.viewTitle = "trace";
    	},
    	// failure
    	function() {
			var overlay = new tr.ui.b.Overlay();
			overlay.textContent = tr.b.normalizeException(err).message;
			overlay.title = 'Import error';
			overlay.visible = true;
    	}
	);
  });
}());
</script>
</head>
<body>
</body>
</html>
`
)
