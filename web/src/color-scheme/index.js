export class NHIEColorScheme {
    constructor() {
        const toggle = document.querySelector('#js-invert-color-scheme__toggle');
        const layout = document.querySelector('.nhie-layout');

        this.rootElement = layout;

        toggle.addEventListener('click', () => {
            layout.classList.toggle('nhie-theme--invert');
        });
    }

    get invertColorScheme() {
        return this.rootElement.classList.contains('nhie-theme--invert');
    }
}