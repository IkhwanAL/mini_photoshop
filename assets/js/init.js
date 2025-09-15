document.addEventListener("DOMContentLoaded", () => {
  initCanvas("remoteCanvas")
})

function resizeCanvas(canvas, ctx = null) {
  const rect = canvas.getBoundingClientRect()
  canvas.width = rect.width
  canvas.height = rect.height

  if (ctx) {
    ctx.fillStyle = "white"
    ctx.fillRect(0,0, canvas.width, canvas.height)
  }
}

function initCanvas(id) {
  /**
   * @type {HTMLCanvasElement}
   */
  const startCanvas = document.getElementById(id)
  if (!startCanvas) {
    return
  }

  resizeCanvas(startCanvas)

  /**
   * @type {HTMLCanvasElement}
   */
  const rulerTop = document.getElementById("rulerTop")
  if (!rulerTop) {
    return
  }

  resizeCanvas(rulerTop)
  /**
   * @type {HTMLCanvasElement}
   */
  const rulerLeft = document.getElementById("rulerLeft")
  if (!rulerLeft) {
    return
  }

  resizeCanvas(rulerLeft)
}

