const { events } = require("brigadier")

events.on("exec", (e, p) => {
  console.log("Hello from brig!")
})
