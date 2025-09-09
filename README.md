# Art Player

## What is this project all about?
Since i already i have knowledge on how to scale, filter image, now it's time to integrated with a web and display it in the web

## what the Motivation of this project?
I do have interest to replicate a figma or other similart web, but try to reverse engineer the web is quite difficult in my head.
so right now just a simple image manipulation is enough. And let's see in the future if it's possible of doing reverse engineer figma.

I've change my mind right now i create a photo manipulation like mini photoshop, but for now let's make it simple so it won't became to complicated to develop. 

## Stack
- Go
- HTMX
- JS
- Templ (Templating for Go)

## Step By Step Journey
For now i start to build something simple like upload an image an able to edit image that's all. 
No need to create complicated UI just upload and canvas to show / edit image. Few days manage to create upload image and paint the image to canvas.

After a few days i did manage to make canvas Zoom in / Out. By using pure Javascript with a little bit a help from AI, and some human intervention manage to built with only 30 - 60 line of code (including newline). Now i want to make zoom like figma or Photoshop where i can drag the image, it will take some time, but let see.

Successfully Built Zoom and Pan so i, after that i can visual what my project will look like, after that i added a database using sqlite to make it simple than add 2 table one table is upload image other is image tracker. upload image is quite straigth forward to store the file header but for tracker image it's the same store file header but every action user made it will be store in tracker image the reason is, when user want to undo or redo i just need to change the pointer of file to previous image.


