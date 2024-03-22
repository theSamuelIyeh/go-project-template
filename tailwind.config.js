/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./*/*.{templ,html}"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],

  daisyui: {
    themes: ["black", "light", "light"],
  },
};
