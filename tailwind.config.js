/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./internal/view/**/*.templ"],
    theme: {
        colors: {
            'text': '#090906',
            'background': '#fdfdfc',
            'primary': '#998d70',
            'secondary': '#d3c39c',
            'accent': '#90c072',
        },
        fontFamily: {
            mono: ['Inconsolata', 'monospace']
        },
        extend: {},
    },
    plugins: [],
}

