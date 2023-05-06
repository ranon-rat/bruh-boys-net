const canvas = document.getElementById("title") as HTMLCanvasElement;
const howMany = 10
const randomPhrases = ["The CIA is after me",
  "Accept your malice",
  "The industrial society and its consequences have been a disaster for the human race",
  "You are immune to propaganda",
  "chess",
  "brave new world", "im not even human anymore", "cryptography", "they have access to my network", "we worship a different god", "I have become death, the destroyer of worlds",
  "they controll us",
  "psyop psyop psyop psyop", "i hate cities", "im in a watchlist", "the key for god is mathematics", "you will own nothing and you will be happy", "👁", "👁", "👁", "👁", "👁", "👁", "👁", "👁", "👁", "👁", "👁", "👁"];
const howManyImages = 29
const ctx = canvas.getContext("2d")!;
// background and canvas style setted in here
canvas.height = canvas.clientHeight;
canvas.width = canvas.clientWidth;
canvas.style.backgroundPosition = "center"
canvas.style.backgroundAttachment = "fixed"
canvas.style.backgroundImage = 'url("static/images/television-static.gif")'

let change = false;

window.addEventListener("resize", () => {
  canvas.height = canvas.clientHeight;
  canvas.width = canvas.clientWidth;
});
function drawText(x: number, y: number, move: number, text: string) {
  let alpha = Math.random() * Math.PI * 2
  let dx = Math.cos(alpha)
  let dy = Math.sin(alpha)
  for (let i = 0; i < 4; i++) {
    ctx.fillStyle = "rgb(" + Math.random() * 200 + 75 + "," + Math.random() * 200 + 75 + "," + Math.random() * 200 + 75 + ")"
    // just to do some rotations, and some cool stuff with it
    ctx.fillText(text, x - move * dx * i + move / 2, y - move * dy * i + move / 2)
    dx = Math.cos(alpha * i)
    dy = Math.sin(alpha * i)
  }


}
function animation(ctx: CanvasRenderingContext2D, canvas: HTMLCanvasElement) {

  ctx.textAlign = "center";
  ctx.font = "2vw sans-serif";
  if (!change) {
    for (let i = 0; i < howMany; i++) {
      ctx.font = "2vw Fantasia";

      let rand = Math.floor(randomPhrases.length * Math.random())// i just select a random phrase
      let phrase = randomPhrases[rand]
      drawText(canvas.width * Math.random(), canvas.height * Math.random(), 1, phrase);
      if (Math.random() < 0.5) {
        ctx.font = Math.random() * 10 + "vw sans-serif";

        drawText(canvas.width * Math.random(), canvas.height * Math.random(), 4, "👁");

      }
    }
  }
  let move = 5;
  // i cant just put the font size without the font so i set this
  // i dont really care but you need to take that in mind
  ctx.font = "10vw Star Wars";
  drawText(canvas.width / 2, canvas.height / 2, move, "👁 Tecnopsicosis 👁");
// this is for keeping something constant
  ctx.fillStyle = "rgb(255,0,0)"
  ctx.fillText("👁 Tecnopsicosis 👁", canvas.width / 2, canvas.height / 2)
}
setInterval(async () => {
  // i need to clear everything

  ctx.clearRect(0, 0, canvas.width, canvas.height);

  requestAnimationFrame(
    animation.bind(animation, ctx, canvas)
  );
}, 100);

setInterval(async () => {

  let rand = Math.round(Math.random() * howManyImages)
  if (!change) {
    // i just select a random gif
    canvas.style.backgroundImage = 'url("static/images/' + rand + '.gif")'
  } else {
    canvas.style.backgroundImage = 'url("static/images/television-static.gif")'

  }

  change = !change;
}, 1500);