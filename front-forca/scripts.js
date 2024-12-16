// var lineDegola = document.querySelector("#meu-svg")
// var svgContent = document.querySelector("#lineDegola")
// console.log(lineDegola)

const word = document.getElementById("word")
const form = document
  .getElementById("form")
  .addEventListener("submit", tryLetter)
const alertPlaceholder = document.getElementById("liveAlert")
const letterDigited = document.getElementById("lettersDigited")
let letter = document.getElementById("letterInput")
let letters = []

async function getWord() {
  try {
    const response = await fetch("http://localhost:1323/word", {
      method: "GET",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
    })

    word.innerHTML = (await response.text()).slice(5)
    // const word = await response.text()
  } catch (error) {
    console.log(error)
  }
}

async function resetGame() {
  try {
    const response = await fetch("http://localhost:1323/word/new", {
      method: "GET",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
    })

    word.innerHTML = (await response.text()).slice(5)
    letter.value = ""
    letters = []
    letterDigited.innerText = letters
  } catch (error) {
    console.log(error)
  }
}

async function tryLetter(event) {
  event.preventDefault()

  if (!letters.includes(letter.value, 0)) {
    letters.push(letter.value)
  }

  console.log(letter.value)

  try {
    const response = await fetch("http://localhost:1323/word", {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        letter: letter.value,
      }),
    })

    const wordChange = await response.text()

    if (!wordChange.includes("_", 0)) {
      // showAlert(wordChange, "success")
      // alert(wordChange)
      word.innerHTML = wordChange
      letterDigited.value = ""
    } else {
      word.innerHTML = wordChange.slice(5)
      letter.value = ""
      letterDigited.innerHTML = letters
      console.log(letters)
    }
  } catch (error) {
    console.log(error)
  }
}

function showAlert(message, type) {
  const wrapper = document.createElement("div")
  wrapper.innerHTML = [
    `<div class="alert alert-${type} alert-dismissible w-100 h-100" role="alert">
      <div>${message}</div>
      <button
        type="button"
        class="btn-close"
        data-bs-dismiss="alert"
        aria-label="Close"
      ></button>
    </div>`,
  ].join("")

  alertPlaceholder.append(wrapper)
}

getWord()
