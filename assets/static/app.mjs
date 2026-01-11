import "https://esm.sh/altcha@2.3.0/es2022/altcha.mjs";

document.querySelector("altcha-widget")?.addEventListener("statechange", (e) => {
	console.log(e);
});
