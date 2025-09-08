/**
 * @type {HTMLCanvasElement}
 */
const canvas = document.getElementById("remoteCanvas")
const ctx = canvas.getContext("2d")

let canvasState = {
  zoom: 1,
  offsetX: 0,
  offsetY: 0,
  isDragging: false,
  image: null,
  lastXPosition: 0, // Last Mouse Position Recorded When Pressing Mouse
  lastYPosition: 0, // Last Mouse Position Recorded When Pressing Mouse
}

// HTMX-Js Will Listen Any Header For Go With HX-Trigger
document.body.addEventListener("updateCanvas", evt => {
  const image = new Image()

  image.onload = ev => {
    ctx.clearRect(0, 0, canvas.width, canvas.height)
    ctx.drawImage(image, 0, 0)
  }
  image.src = evt.detail.image

  canvasState.image = image
})


function draw() {
  const {zoom, offsetX, offsetY} = canvasState
  ctx.setTransform(zoom, 0, 0, zoom, offsetX, offsetY) // Zoom Part
  
  ctx.clearRect(-offsetX/zoom, -offsetY/zoom, canvas.width / zoom, canvas.height / zoom) // Clear The Previous Img State

  ctx.drawImage(canvasState.image, 0, 0)
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

  canvasState.offsetX += deltaX
  canvasState.offsetY += deltaY

  canvasState.lastXPosition = ev.clientX
  canvasState.lastYPosition = ev.clientY

  draw()
})

canvas.addEventListener("mouseup", () => canvasState.isDragging = false)
canvas.addEventListener("mouseleave", () => canvasState.isDragging = false)

canvas.addEventListener("wheel", ev => {
  console.log("Scroll")

  if (ev.deltaY < 0) {
    canvasState.zoom *= 1.1
  }else{
    canvasState.zoom /= 1.1
  }

  draw()
})
