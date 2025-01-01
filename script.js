const body = document.body

const btnTheme = document.querySelector('.fa-moon')
const btnHamburger = document.querySelector('.fa-bars')

const addThemeClass = (bodyClass, btnClass) => {
  body.classList.add(bodyClass)
  btnTheme.classList.add(btnClass)
}

const getBodyTheme = localStorage.getItem('portfolio-theme') || 'dark'
const getBtnTheme = localStorage.getItem('portfolio-btn-theme') || 'fa-sun'

addThemeClass(getBodyTheme, getBtnTheme)

const isDark = () => body.classList.contains('dark')

const setTheme = (bodyClass, btnClass) => {

  body.classList.remove(localStorage.getItem('portfolio-theme') || 'dark')
  btnTheme.classList.remove(localStorage.getItem('portfolio-btn-theme') || 'fa-sun')

  addThemeClass(bodyClass, btnClass)

  localStorage.setItem('portfolio-theme', bodyClass)
  localStorage.setItem('portfolio-btn-theme', btnClass)
}

const toggleTheme = () =>
  isDark() ? setTheme('light', 'fa-moon') : setTheme('dark', 'fa-sun')

btnTheme.addEventListener('click', toggleTheme)

const displayList = () => {
  const navUl = document.querySelector('.nav__list')

  if (btnHamburger.classList.contains('fa-bars')) {
    btnHamburger.classList.remove('fa-bars')
    btnHamburger.classList.add('fa-times')
    navUl.classList.add('display-nav-list')
  } else {
    btnHamburger.classList.remove('fa-times')
    btnHamburger.classList.add('fa-bars')
    navUl.classList.remove('display-nav-list')
  }
}

btnHamburger.addEventListener('click', displayList)

const scrollUp = () => {
  const btnScrollTop = document.querySelector('.scroll-top')

  if (
    body.scrollTop > 500 ||
    document.documentElement.scrollTop > 500
  ) {
    btnScrollTop.style.display = 'block'
  } else {
    btnScrollTop.style.display = 'none'
  }
}

document.addEventListener('scroll', scrollUp)
const titleElement = document.querySelector('.about__role');
const titleText = "An Aspiring Software Engineer.";
let currentIndex = 0;
let isForward = true;

function typeTitle() {
  const currentText = titleElement.textContent;
  if (isForward) {
    if (currentIndex < titleText.length) {
      titleElement.textContent = currentText + titleText[currentIndex];
      currentIndex++;
      setTimeout(typeTitle, 100); // Adjust the typing speed (milliseconds) as needed
    } else {
      isForward = false;
      setTimeout(typeTitle, 1000); // Pause at the end before reversing
    }
  } else {
    if (currentIndex > 0) {
      titleElement.textContent = currentText.slice(0, -1);
      currentIndex--;
      setTimeout(typeTitle, 50); // Adjust the backspacing speed (milliseconds) as needed
    } else {
      isForward = true;
      setTimeout(typeTitle, 1000); // Pause at the beginning before typing forward again
    }
  }
}

// Start the typewriter effect when the page loads
window.addEventListener('load', () => {
  setTimeout(typeTitle, 1000); // Delay before starting the typewriter effect
});
