/**
 * @type {HTMLCanvasElement}
 */
const canvas = document.getElementById("remoteCanvas")
const ctx = canvas.getContext("2d")

class CanvasState extends EventTarget {
  constructor() {
    super()
    this._zoom = 1
    this._offsetX = 0
    this._offsetY = 0
    this._isDragging = false
    this._image = null
    this._lastX = 0
    this._lastY = 0
  }

  get zoom() { return this._zoom }
  set zoom(val) {
    if (this._zoom != val) {
      this._zoom = val
      this.dispatchEvent(new Event("change"))
    }
  }

  get lastX() { return this._lastX }
  set lastX(val) {
    if (this._lastX != val) {
      this._lastX = val
    }
  }
  get lastY() { return this._lastY }
  set lastY(val) {
    if (this._lastY != val) {
      this._lastY = val
    }
  }

  get isDragging() { return this._isDragging }
  set isDragging(val) {
    if (this._isDragging != val) {
      this._isDragging = val
    }
  }  
  
  get offsetX() { return this._offsetX }
  get offsetY() { return this._offsetY }
  
  setOffsetXY = (deltaX, deltaY) => {
    this._offsetX += deltaX
    this._offsetY += deltaY

    this.dispatchEvent(new Event("change"))
  }
}

const canvasState = new CanvasState()

canvasState.addEventListener("change", () => {
  requestAnimationFrame(() => {
    scheduleRedraw()
  })
})

/**
 * @type {AbortController}
 */
let loadController = null

// HTMX-Js Will Listen Any Header For Go With HX-Trigger
document.body.addEventListener("updateCanvas", evt => {
  if (!canvas.classList.contains("block")) {
    canvas.classList.remove("hidden")
    canvas.classList.add("block")
    resizeCanvas(canvas, ctx)
  }

  if (loadController) loadController.abort()
  loadController = new AbortController()

  const image = new Image()
  image.src = evt.detail.image

  image.onload = ev => {
    ctx.clearRect(0, 0, canvas.width, canvas.height)
    ctx.drawImage(image, 0, 0)
    canvasState.image = image
    draw()
  }
})

function draw() {
  const { zoom, offsetX, offsetY } = canvasState
  ctx.setTransform(zoom, 0, 0, zoom, offsetX, offsetY) // Zoom Part

  ctx.clearRect(-offsetX / zoom, -offsetY / zoom, canvas.width / zoom, canvas.height / zoom) // Clear The Previous Img State

  ctx.drawImage(canvasState.image, 0, 0)

  drawRuler()
}

let needsRedraw = false

function scheduleRedraw() {
  if (!needsRedraw) {
    needsRedraw = true
    requestAnimationFrame(() => {
      draw()
      needsRedraw = false
    })
  }
}

// Where Dragging Start
canvas.addEventListener("mousedown", (ev) => {
  canvasState.lastXPosition = ev.clientX
  canvasState.lastYPosition = ev.clientY
  canvasState.isDragging = true
})

canvas.addEventListener("mousemove", (ev) => {
  if (!canvasState.isDragging) return

  const deltaX = ev.clientX - canvasState.lastXPosition
  const deltaY = ev.clientY - canvasState.lastYPosition

  // canvasState.offsetX += deltaX
  // canvasState.offsetY += deltaY

  canvasState.setOffsetXY(deltaX, deltaY)

  canvasState.lastXPosition = ev.clientX
  canvasState.lastYPosition = ev.clientY

})

canvas.addEventListener("mouseup", () => canvasState.isDragging = false)
canvas.addEventListener("mouseleave", () => canvasState.isDragging = false)

canvas.addEventListener("wheel", ev => {

  if (ev.deltaY < 0) {
    canvasState.zoom *= 1.1
  } else {
    canvasState.zoom /= 1.1
  }

})

/**
 * @type {HTMLCanvasElement}
 */
const rulerTop = document.getElementById("rulerTop");
/**
 * @type {HTMLCanvasElement}
 */
const rulerLeft = document.getElementById("rulerLeft");

function drawRuler() {
  const ctxTopRuler = rulerTop.getContext("2d")
  const ctxLeftRuler = rulerLeft.getContext("2d")

  const { zoom, offsetX, offsetY } = canvasState
  const spacing = 50

  ctxTopRuler.clearRect(0, 0, rulerTop.width, rulerTop.height);
  ctxLeftRuler.clearRect(0, 0, rulerLeft.width, rulerLeft.height);

  ctxTopRuler.strokeStyle = ctxLeftRuler.strokeStyle = "#444";
  ctxTopRuler.fillStyle = ctxLeftRuler.fillStyle = "#444";
  ctxTopRuler.font = ctxLeftRuler.font = "10px sans-sarif"

  for (let x = offsetX % (spacing * zoom); x < rulerTop.width; x += spacing * zoom) {
    const worldX = Math.round((x - offsetX) / zoom)
    ctxTopRuler.beginPath()
    ctxTopRuler.moveTo(x, rulerTop.height)
    ctxTopRuler.lineTo(x, rulerTop.height - 5)
    ctxTopRuler.stroke()
    if (worldX % 100 == 0) ctxTopRuler.fillText(worldX, x, 10)
  }
  for (let y = offsetY % (spacing * zoom); y < rulerLeft.height; y += spacing * zoom) {
    const worldY = Math.round((y - offsetY) / zoom)
    ctxLeftRuler.beginPath()
    ctxLeftRuler.moveTo(rulerLeft.width, y)
    ctxLeftRuler.lineTo(rulerLeft.width - 2, y)
    ctxLeftRuler.stroke()
    if (worldY % 100 == 0) ctxLeftRuler.fillText(worldY, 2, y)
  }
}
