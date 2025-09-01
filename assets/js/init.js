document.addEventListener("DOMContentLoaded", () => {
  initCanvas("remoteCanvas")
})

document.body.addEventListener("htmx:afterswap", (evt) => {
  const canvas = evt.target.querySelector("#remoteCanvas");
  console.log("My Canvas", canvas)
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
