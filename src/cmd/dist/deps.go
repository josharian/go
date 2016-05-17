// generated by mkdeps.bash

package main

var builddeps = map[string][]string{
	"bufio":                             {"bytes", "errors", "internal/race", "io", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sync", "sync/atomic", "unicode", "unicode/utf8"},
	"bytes":                             {"errors", "internal/race", "io", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sync", "sync/atomic", "unicode", "unicode/utf8"},
	"cmd/internal/traceviewer":          {"errors", "internal/race", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sync", "sync/atomic", "syscall", "time", "unicode/utf16"},
	"compress/flate":                    {"bufio", "bytes", "errors", "fmt", "internal/race", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "math", "os", "reflect", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sort", "strconv", "sync", "sync/atomic", "syscall", "time", "unicode", "unicode/utf16", "unicode/utf8"},
	"compress/zlib":                     {"bufio", "bytes", "compress/flate", "errors", "fmt", "hash", "hash/adler32", "internal/race", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "math", "os", "reflect", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sort", "strconv", "sync", "sync/atomic", "syscall", "time", "unicode", "unicode/utf16", "unicode/utf8"},
	"container/heap":                    {"runtime", "runtime/internal/atomic", "runtime/internal/sys", "sort"},
	"context":                           {"errors", "fmt", "internal/race", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "math", "os", "reflect", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "strconv", "sync", "sync/atomic", "syscall", "time", "unicode/utf16", "unicode/utf8"},
	"crypto":                            {"errors", "hash", "internal/race", "io", "math", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "strconv", "sync", "sync/atomic", "unicode/utf8"},
	"crypto/sha1":                       {"crypto", "errors", "hash", "internal/race", "io", "math", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "strconv", "sync", "sync/atomic", "unicode/utf8"},
	"debug/dwarf":                       {"encoding/binary", "errors", "fmt", "internal/race", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "math", "os", "path", "reflect", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sort", "strconv", "strings", "sync", "sync/atomic", "syscall", "time", "unicode", "unicode/utf16", "unicode/utf8"},
	"debug/elf":                         {"bufio", "bytes", "compress/flate", "compress/zlib", "debug/dwarf", "encoding/binary", "errors", "fmt", "hash", "hash/adler32", "internal/race", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "math", "os", "path", "reflect", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sort", "strconv", "strings", "sync", "sync/atomic", "syscall", "time", "unicode", "unicode/utf16", "unicode/utf8"},
	"debug/macho":                       {"bytes", "debug/dwarf", "encoding/binary", "errors", "fmt", "internal/race", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "math", "os", "path", "reflect", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sort", "strconv", "strings", "sync", "sync/atomic", "syscall", "time", "unicode", "unicode/utf16", "unicode/utf8"},
	"encoding":                          {"runtime", "runtime/internal/atomic", "runtime/internal/sys"},
	"encoding/base64":                   {"errors", "internal/race", "io", "math", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "strconv", "sync", "sync/atomic", "unicode/utf8"},
	"encoding/binary":                   {"errors", "internal/race", "io", "math", "reflect", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "strconv", "sync", "sync/atomic", "unicode/utf8"},
	"encoding/json":                     {"bytes", "encoding", "encoding/base64", "errors", "fmt", "internal/race", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "math", "os", "reflect", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sort", "strconv", "strings", "sync", "sync/atomic", "syscall", "time", "unicode", "unicode/utf16", "unicode/utf8"},
	"errors":                            {"runtime", "runtime/internal/atomic", "runtime/internal/sys"},
	"flag":                              {"errors", "fmt", "internal/race", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "math", "os", "reflect", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sort", "strconv", "sync", "sync/atomic", "syscall", "time", "unicode/utf16", "unicode/utf8"},
	"fmt":                               {"errors", "internal/race", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "math", "os", "reflect", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "strconv", "sync", "sync/atomic", "syscall", "time", "unicode/utf16", "unicode/utf8"},
	"go/ast":                            {"bytes", "errors", "fmt", "go/scanner", "go/token", "internal/race", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "math", "os", "path/filepath", "reflect", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sort", "strconv", "strings", "sync", "sync/atomic", "syscall", "time", "unicode", "unicode/utf16", "unicode/utf8"},
	"go/build":                          {"bufio", "bytes", "errors", "fmt", "go/ast", "go/doc", "go/parser", "go/scanner", "go/token", "internal/race", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "io/ioutil", "log", "math", "net/url", "os", "path", "path/filepath", "reflect", "regexp", "regexp/syntax", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sort", "strconv", "strings", "sync", "sync/atomic", "syscall", "text/template", "text/template/parse", "time", "unicode", "unicode/utf16", "unicode/utf8"},
	"go/doc":                            {"bytes", "errors", "fmt", "go/ast", "go/scanner", "go/token", "internal/race", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "io/ioutil", "math", "net/url", "os", "path", "path/filepath", "reflect", "regexp", "regexp/syntax", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sort", "strconv", "strings", "sync", "sync/atomic", "syscall", "text/template", "text/template/parse", "time", "unicode", "unicode/utf16", "unicode/utf8"},
	"go/parser":                         {"bytes", "errors", "fmt", "go/ast", "go/scanner", "go/token", "internal/race", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "io/ioutil", "math", "os", "path/filepath", "reflect", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sort", "strconv", "strings", "sync", "sync/atomic", "syscall", "time", "unicode", "unicode/utf16", "unicode/utf8"},
	"go/scanner":                        {"bytes", "errors", "fmt", "go/token", "internal/race", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "math", "os", "path/filepath", "reflect", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sort", "strconv", "strings", "sync", "sync/atomic", "syscall", "time", "unicode", "unicode/utf16", "unicode/utf8"},
	"go/token":                          {"errors", "fmt", "internal/race", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "math", "os", "reflect", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sort", "strconv", "sync", "sync/atomic", "syscall", "time", "unicode/utf16", "unicode/utf8"},
	"hash":                              {"errors", "internal/race", "io", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sync", "sync/atomic"},
	"hash/adler32":                      {"errors", "hash", "internal/race", "io", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sync", "sync/atomic"},
	"internal/race":                     {"runtime", "runtime/internal/atomic", "runtime/internal/sys"},
	"internal/singleflight":             {"internal/race", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sync", "sync/atomic"},
	"internal/syscall/windows":          {"errors", "internal/race", "internal/syscall/windows/sysdll", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sync", "sync/atomic", "syscall", "unicode/utf16"},
	"internal/syscall/windows/registry": {"errors", "internal/race", "internal/syscall/windows/sysdll", "io", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sync", "sync/atomic", "syscall", "unicode/utf16"},
	"internal/syscall/windows/sysdll":   {"runtime", "runtime/internal/atomic", "runtime/internal/sys"},
	"io":                      {"errors", "internal/race", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sync", "sync/atomic"},
	"io/ioutil":               {"bytes", "errors", "internal/race", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "math", "os", "path/filepath", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sort", "strconv", "strings", "sync", "sync/atomic", "syscall", "time", "unicode", "unicode/utf16", "unicode/utf8"},
	"log":                     {"errors", "fmt", "internal/race", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "math", "os", "reflect", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "strconv", "sync", "sync/atomic", "syscall", "time", "unicode/utf16", "unicode/utf8"},
	"math":                    {"runtime", "runtime/internal/atomic", "runtime/internal/sys"},
	"net/url":                 {"bytes", "errors", "fmt", "internal/race", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "math", "os", "reflect", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sort", "strconv", "strings", "sync", "sync/atomic", "syscall", "time", "unicode", "unicode/utf16", "unicode/utf8"},
	"os":                      {"errors", "internal/race", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sync", "sync/atomic", "syscall", "time", "unicode/utf16", "unicode/utf8"},
	"os/exec":                 {"bytes", "context", "errors", "fmt", "internal/race", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "math", "os", "path/filepath", "reflect", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sort", "strconv", "strings", "sync", "sync/atomic", "syscall", "time", "unicode", "unicode/utf16", "unicode/utf8"},
	"os/signal":               {"errors", "internal/race", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "os", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sync", "sync/atomic", "syscall", "time", "unicode/utf16", "unicode/utf8"},
	"path":                    {"errors", "internal/race", "io", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "strings", "sync", "sync/atomic", "unicode", "unicode/utf8"},
	"path/filepath":           {"errors", "internal/race", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "os", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sort", "strings", "sync", "sync/atomic", "syscall", "time", "unicode", "unicode/utf16", "unicode/utf8"},
	"reflect":                 {"errors", "internal/race", "math", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "strconv", "sync", "sync/atomic", "unicode/utf8"},
	"regexp":                  {"bytes", "errors", "internal/race", "io", "math", "regexp/syntax", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sort", "strconv", "strings", "sync", "sync/atomic", "unicode", "unicode/utf8"},
	"regexp/syntax":           {"bytes", "errors", "internal/race", "io", "math", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sort", "strconv", "strings", "sync", "sync/atomic", "unicode", "unicode/utf8"},
	"runtime":                 {"runtime/internal/atomic", "runtime/internal/sys"},
	"runtime/internal/atomic": {"runtime/internal/sys"},
	"runtime/internal/sys":    {},
	"sort":                    {"runtime", "runtime/internal/atomic", "runtime/internal/sys"},
	"strconv":                 {"errors", "math", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "unicode/utf8"},
	"strings":                 {"errors", "internal/race", "io", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sync", "sync/atomic", "unicode", "unicode/utf8"},
	"sync":                    {"internal/race", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sync/atomic"},
	"sync/atomic":             {"runtime", "runtime/internal/atomic", "runtime/internal/sys"},
	"syscall":                 {"errors", "internal/race", "internal/syscall/windows/sysdll", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sync", "sync/atomic", "unicode/utf16"},
	"text/template":           {"bytes", "errors", "fmt", "internal/race", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "io/ioutil", "math", "net/url", "os", "path/filepath", "reflect", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sort", "strconv", "strings", "sync", "sync/atomic", "syscall", "text/template/parse", "time", "unicode", "unicode/utf16", "unicode/utf8"},
	"text/template/parse":     {"bytes", "errors", "fmt", "internal/race", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "math", "os", "reflect", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "strconv", "strings", "sync", "sync/atomic", "syscall", "time", "unicode", "unicode/utf16", "unicode/utf8"},
	"time":                    {"errors", "internal/race", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sync", "sync/atomic", "syscall", "unicode/utf16"},
	"unicode":                 {"runtime", "runtime/internal/atomic", "runtime/internal/sys"},
	"unicode/utf16":           {"runtime", "runtime/internal/atomic", "runtime/internal/sys"},
	"unicode/utf8":            {"runtime", "runtime/internal/atomic", "runtime/internal/sys"},
	"cmd/go":                  {"bufio", "bytes", "cmd/internal/traceviewer", "compress/flate", "compress/zlib", "container/heap", "context", "crypto", "crypto/sha1", "debug/dwarf", "debug/elf", "debug/macho", "encoding", "encoding/base64", "encoding/binary", "encoding/json", "errors", "flag", "fmt", "go/ast", "go/build", "go/doc", "go/parser", "go/scanner", "go/token", "hash", "hash/adler32", "internal/race", "internal/singleflight", "internal/syscall/windows", "internal/syscall/windows/registry", "internal/syscall/windows/sysdll", "io", "io/ioutil", "log", "math", "net/url", "os", "os/exec", "os/signal", "path", "path/filepath", "reflect", "regexp", "regexp/syntax", "runtime", "runtime/internal/atomic", "runtime/internal/sys", "sort", "strconv", "strings", "sync", "sync/atomic", "syscall", "text/template", "text/template/parse", "time", "unicode", "unicode/utf16", "unicode/utf8"},
}
