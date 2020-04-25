(function () {
    var codemirror = CodeMirror(document.body, {
        value: '// CodeMirror Addon hint/show-hint.js sample.\n// Snippets are Ctrl-E or Cmd-E.',
        mode: 'text/javascript',
        lineNumbers: true,
        styleActiveLine: true,
        theme: 'solarized dark'
    });
    // keymap
    codemirror.setOption('extraKeys', {
        'Cmd-E': function () {
            snippet();
        },
        'Ctrl-E': function () {
            snippet();
        },
        'Ctrl-,': function () {
            snippet();
        },
        'Ctrl-.': function () {
            snippet();
        },
        'Ctrl-/': function () {
            snippet();
        }
    });
    var snippets = [
        { text: 'insert ""', displayText: 'insert: add contents of specified template.' },
        { text: 'nid ""', displayText: 'nid: no id. Insert text as is.' },
        { text: 'rid ".*~"', displayText: 'rid: insert text using relative ID.' },
        { text: 'sid "~"', displayText: 'sid: insert text using specific ID.' },
        { text: 'ver "~"', displayText: 'ver: display version name.' },
        { text: 'when_exists rid ".*~" use:  {\n\totherwise use:\n}', displayText: 'when_exists' },
    ];
    function snippet() {
        CodeMirror.showHint(codemirror, function () {
            var cursor = codemirror.getCursor();
            var token = codemirror.getTokenAt(cursor);
            var start = token.start;
            var end = cursor.ch;
            var line = cursor.line;
            var currentWord = token.string;
            var list = snippets.filter(function (item) {
                return item.text.indexOf(currentWord) >= 0;
            });
            return {
                list: list.length ? list : snippets,
                from: CodeMirror.Pos(line, start),
                to: CodeMirror.Pos(line, end)
            };
        }, { completeSingle: false });
    }
})();
