import './style.css';

import {NHIECategory} from './src/category';
import {NHIECategorySelection} from './src/category-selection';
import {NHIEClient} from './src/client';
import {NHIEColorScheme} from "./src/color-scheme";
import {NHIELanguageSelection} from './src/language-selection';
import {NHIELinearProgress} from "./src/linear-progress";
import {NHIEStatement} from './src/statement';

document.addEventListener('DOMContentLoaded', () => {
  limitSearchParams('language');

  const statement = new NHIEStatement();

  const categories = [].slice.call(document.querySelectorAll('.nhie-category__input')).map((e) => new NHIECategory(e));

  const categorySelection = new NHIECategorySelection();
  const languageSelection = new NHIELanguageSelection();

  const linearProgress = new NHIELinearProgress();
  const colorScheme = new NHIEColorScheme();

  const client = new NHIEClient();

  categorySelection.registerCategories(...categories);

  client.registerFeature('statement_id', () => statement.ID);
  client.registerFeature('category', () => categorySelection.activeCategories.map((c) => c.key));
  client.registerFeature('language', () => languageSelection.activeLanguage);
  client.registerFeature('invert_color_scheme', () => colorScheme.invertColorScheme ? 'true' : 'false');

  document.addEventListener('keyup', (keyboardEvent) => {
    if (keyboardEvent.code === 'Space') {
      refreshStatement();
    }
  });

  document.querySelector('main').addEventListener('click', () => {
    refreshStatement();
  });

  document.addEventListener('language-changed', () => {
    let url = new URL(window.location.href);
    client.encodeFeatures(url, client.limitFeatures('language', 'category', 'invert_color_scheme'));
    window.location.assign(url);
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
          replaceStatementID(result.ID);
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

function limitSearchParams() {
  let url = new URL(window.location.href);
  let keys = Array.from(arguments);

  let removals = Array();

  url.searchParams.forEach((value, key, ) => {
    if (keys.includes(key)) {
      return;
    }

    if (removals.includes(key)) {
      return;
    }

    removals.push(key);
  });

  removals.forEach((key) => {url.searchParams.delete(key)})

  window.history.replaceState({}, document.title, url);
}

function replaceStatementID(ID) {
  let url = new URL(window.location.href);
  let path = url.pathname.split('/');

  path[path.length - 1] = ID;
  url.pathname = path.join('/');

  window.history.replaceState({}, document.title, url);
}
