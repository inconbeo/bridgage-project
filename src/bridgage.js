const { events } = require("brigadier")

events.on("exec", (brigadeEvent, project) => {
  console.log("Hello from brig!")
})
