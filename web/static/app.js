document.addEventListener('DOMContentLoaded', (event) => {
    const form = document.querySelector('#js-form');
    const statementID = document.querySelector('#js-statement-id').getAttribute('value');

    initLanguageSelection(() => {
        form.setAttribute('action', '/statements/' + statementID);
        form.submit();
    });

    initCategorySelection((categories) => {
        if (categories.filter((category) => category.checked).length === 0) {
            categories.forEach((category) => {category.checked = true});
        }
    });

    document.querySelector('main').addEventListener('click', () => {
        form.submit();
    });

    document.addEventListener('keydown', (keyboardEvent) => {
        if (keyboardEvent.code === 'Space') {
            form.submit();
        }
    });
});

function initLanguageSelection(callback) {
    const toggle = document.querySelector('#js-language-selection__toggle');
    const selection = document.querySelector('#js-language-selection');

    toggle.addEventListener('click', (event) => {
        selection.classList.toggle('nhie-language-selection--hidden');
    });

    document.querySelectorAll('.nhie-language-selection__input').forEach((language) => {
        language.addEventListener('click', (event) => {
            selection.classList.toggle('nhie-language-selection--hidden');
            document.querySelector('#js-statement-id').setAttribute('disabled', '');

            callback();
        });
    });

}

function initCategorySelection(callback) {
    document.querySelector('#js-category-selection').addEventListener('click', (event) => {
        event.stopPropagation();
    });

    let categories = [];

    document.querySelectorAll('.nhie-category__input').forEach((category) => {
        categories.push(category)

        category.addEventListener('change', () => {
            callback(categories);
        });
    });
}
