const filterOption = document.getElementById("filter")
const compBrightnessContrast = document.getElementById("brightnessContrast")

filterOption.addEventListener("change", (ev) => {
	if (ev.target.value == "brightness-contrast") {
		compBrightnessContrast.classList.remove("hidden")
	}else{
		compBrightnessContrast.classList.add("hidden")
	}
	
})

const brightness = document.getElementById("brightness")
const brightnessValue = document.getElementById("brightnessValue")

brightness.addEventListener("input", ev => {
	brightnessValue.textContent = ev.target.value
})

const contrast = document.getElementById("contrast")
const contrastValue = document.getElementById("contrastValue")

contrast.addEventListener("input", ev => {
	contrastValue.textContent = ev.target.value
})

