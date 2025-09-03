document.addEventListener("DOMContentLoaded", () => {
  initCanvas("remoteCanvas")
})

document.body.addEventListener("htmx:afterswap", (evt) => {
  const canvas = evt.target.querySelector("#remoteCanvas");
  if (canvas) {
    initCanvas("remoteCanvas");
  }
});

function initCanvas(id) {
  /**
   * @type {HTMLCanvasElement}
   */
  const canvas = document.getElementById(id)

  if (!canvas) {
    return
  }

  const ctx = canvas.getContext("2d")

  function resizeCanvas() {
    const rect = canvas.getBoundingClientRect()
    canvas.width = rect.width
    canvas.height = rect.height

    ctx.fillStyle = "white"
    ctx.fillRect(0, 0, canvas.width, canvas.height)
  }

  resizeCanvas()
}

// HTMX-Js Will Listen Any Header For Go With HX-Trigger
document.body.addEventListener("updateCanvas", evt => {
   console.log(evt.detail)
  /**
   * @type {HTMLCanvasElement}
   */
  const canvas = document.getElementById("remoteCanvas")
  const ctx = canvas.getContext("2d")

  const image = new Image()

  image.onload = ev => {
    ctx.clearRect(0,0, canvas.width, canvas.height)
    ctx.drawImage(image, 0,0)
  }
  image.src = evt.detail.image
})


