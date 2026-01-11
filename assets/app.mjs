import "https://esm.sh/altcha@2.3.0";

window.addEventListener("load", () => {
	document.querySelector("#altcha")?.configure({
		strings: {
			label: "I am not a pesky clanker",
		},
	});
});
