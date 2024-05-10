/** @type {import('tailwindcss').Config} */
const plugin = require('tailwindcss/plugin')

module.exports = {
    content: ["./internal/view/**/*.templ"],
    theme: {
        fontFamily: {
            mono: ['Inconsolata', 'monospace']
        },
        extend: {
            colors: {
                'text': '#090906',
                'background': '#fdfdfc',
                'primary': '#998d70',
                'secondary': '#d3c39c',
                'accent': '#90c072',
            },
        },
    },
    plugins: [
        plugin(function({ addVariant }) {
          addVariant('htmx-settling', ['&.htmx-settling', '.htmx-settling &'])
          addVariant('htmx-request',  ['&.htmx-request',  '.htmx-request &'])
          addVariant('htmx-swapping', ['&.htmx-swapping', '.htmx-swapping &'])
          addVariant('htmx-added',    ['&.htmx-added',    '.htmx-added &'])
        }),
        require('@tailwindcss/forms')
    ],
}

