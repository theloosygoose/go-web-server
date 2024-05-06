HandleBlur()
htmx.on("htmx:load", function(){
    HandleBlur()
});

function HandleBlur(){
    let blurImg = document.querySelectorAll(".blur-load")

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
}




