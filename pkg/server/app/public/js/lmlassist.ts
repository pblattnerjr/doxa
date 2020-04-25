(function () {

    const codemirror = CodeMirror(document.body, {
        value: '// CodeMirror Addon hint/show-hint.js sample.\n// Snippets are Ctrl-E or Cmd-E.',
        mode: 'text/javascript',
        lineNumbers: true,
        styleActiveLine: true,
        theme: 'solarized dark'
    })

    // keymap
    codemirror.setOption('extraKeys', {
        'Cmd-E': function() {
            snippet()
        },
        'Ctrl-E': function() {
            snippet()
        },
        'Ctrl-,': function() {
            snippet()
        },
        'Ctrl-.': function() {
            snippet()
        },
        'Ctrl-/': function() {
            snippet()
        }
    })

    const snippets: any[] = [
        { text: 'insert ""', displayText: 'insert: add contents of specified template.' },
        { text: 'nid ""', displayText: 'nid: no id. Insert text as is.' },
        { text: 'rid ".*~"', displayText: 'rid: insert text using relative ID.' },
        { text: 'sid "~"', displayText: 'sid: insert text using specific ID.' },
        { text: 'ver "~"', displayText: 'ver: display version name.' },
        { text: 'when_exists rid ".*~" use:  {\n\totherwise use:\n}', displayText: 'when_exists' },
    ]
    function snippet(): void {
        CodeMirror.showHint(codemirror, function (): any {
            const cursor = codemirror.getCursor()
            const token = codemirror.getTokenAt(cursor)
            const start: number = token.start
            const end: number = cursor.ch
            const line: number = cursor.line
            const currentWord: string = token.string

            const list: any[] = snippets.filter(function (item): boolean {
                return item.text.indexOf(currentWord) >= 0
            })

            return {
                list: list.length ? list : snippets,
                from: CodeMirror.Pos(line, start),
                to: CodeMirror.Pos(line, end)
            }
        }, { completeSingle: false })
    }
})()

