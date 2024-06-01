HandleBlur()
htmx.on("htmx:afterSettle", function(){
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
};


var loadFile = function(event) {
    var output = document.getElementById('preview-image');
    output.src = URL.createObjectURL(event.target.files[0]);
    output.onload =function () {
        URL.revokeObjectURL(output.src)
    }
};

document.body.addEventListener('htmx:afterSwap', function(e){
    const form = document.querySelector('#form');
    form.reset();
})

document.getElementById('main-image-photo').addEventListener("click", showDetails);

function showDetails(){
    console.log("Triggered")
    var eclass = document.getElementById('main-details').classList
    if (eclass.contains("details-show")){
        eclass.remove("details-show")
        eclass.add("details-hide")

    } else if (eclass.contains("details-hide")){
        eclass.remove("details-hide")
        eclass.add("details-show")
    }
}
