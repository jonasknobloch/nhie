export class NHIEColorScheme {
    constructor() {
        const toggle = document.querySelector('#js-invert-color-scheme__toggle');
        const layout = document.querySelector('.nhie-layout');

        toggle.addEventListener('click', () => {
            layout.classList.toggle('nhie-theme--invert');

            console.log('foobar');
        });
    }
}