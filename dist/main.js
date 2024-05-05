const blurredImageDiv = document.querySelector(".blur-load")

const img = blurredImageDiv.querySelector("img")

function loaded() {
    blurredImageDiv.classList.add("loaded")
}

if (img.complete) {
    loaded()
} else {
    img.addEventListener("load", loaded)
}
