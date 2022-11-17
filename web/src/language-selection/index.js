import './style.css';

export class NHIELanguageSelection {
  constructor() {
    const toggle = document.querySelector('#js-language-selection__toggle');
    const selection = document.querySelector('#js-language-selection');

    const event = new Event('language-changed');

    toggle.addEventListener('click', () => {
      selection.classList.toggle('nhie-language-selection--hidden');
    });

    document.querySelectorAll('.nhie-language-selection__input').forEach((language) => {
      language.addEventListener('click', () => {
        selection.classList.toggle('nhie-language-selection--hidden');
        document.dispatchEvent(event);
      });
    });
  }

  get activeLanguage() {
    return [].slice.call(document.getElementsByName('language')).find((radio) => radio.checked).value;
  }
}
