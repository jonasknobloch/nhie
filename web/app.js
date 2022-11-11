import './style.css';

import {NHIECategory} from './src/category';
import {NHIECategorySelection} from './src/category-selection';
import {NHIEClient} from './src/client';
import {NHIELanguageSelection} from './src/language-selection';
import {NHIELinearProgress} from "./src/linear-progress";
import {NHIEStatement} from './src/statement';

document.addEventListener('DOMContentLoaded', () => {
    const statement = new NHIEStatement();

    const categories = [].slice.call(document.querySelectorAll('.nhie-category__input')).map((e) => new NHIECategory(e));

    const categorySelection = new NHIECategorySelection();
    const languageSelection = new NHIELanguageSelection();

    const linearProgress = new NHIELinearProgress();

    const client = new NHIEClient();

    categorySelection.registerCategories(...categories);

    client.registerFeature('statement_id', () => statement.ID);
    client.registerFeature('category', () => categorySelection.activeCategories.map((c) => c.key));
    client.registerFeature('language', () => languageSelection.activeLanguage);

    document.addEventListener('keyup', (keyboardEvent) => {
        if (keyboardEvent.code === 'Space') {
            refreshStatement();
        }
    });

    document.addEventListener('language-changed', () => {
        refreshStatement();
    });

    document.querySelector('main').addEventListener('click', () => {
        refreshStatement();
    });

    let refreshing = false;

    function refreshStatement() {
        if (!refreshing) {
            refreshing = true;
            linearProgress.on();

            client.fetchStatement()
                .then((response) => {
                    return response.json()
                })
                .then((result) => {
                    let url = new URL(window.location.href);
                    let path = url.pathname.split('/');

                    path[path.length - 1] = result.ID;
                    url.pathname = path.join('/');

                    window.history.replaceState({}, document.title, url.toString());

                    return result;
                })
                .then((result) => {
                    statement.ID = result.ID;
                    statement.statement = result.statement;
                })
                .catch((error) => {
                    console.log(error);
                    statement.statement = 'Something went wrong.';
                })
                .finally(() => {
                    refreshing = false;
                    linearProgress.off();
                });
        }
    }
});