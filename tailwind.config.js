/** @type {import('tailwindcss').Config} */
export const content = [
	"./templates/**/*.templ",
	"./**/*.html",
	"./**/*.templ",
	"./**/*.go",
];
export const theme = {
	extend: {
		fontFamily: {
			mono: ["Courier Prime", "monospace"],
		},
	},
};
export const plugins = [];
export const corePlugins = { preFlight: true };
