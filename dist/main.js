const blurImg = document.querySelectorAll(".blur-load")
console.log(blurImg)

blurImg.forEach(div => {
    const img = div.querySelector("img")

    function loaded() {
        div.classList.add("loaded")
    }

    if (img.complete){
        loaded()
    } else{
        img.addEventListener("load", loaded)
    }
})



