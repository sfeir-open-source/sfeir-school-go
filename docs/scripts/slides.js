// allFiles : canonical structure for ordered 2-levels filesystem tree
// array of [directory, [file]] pairs

const go100 = [
  ['00-school', ['00-title', '01-speakers']],
  ['01-intro', ['00-title', '01-Le_but_de_go']],
  ['02-installation', ['00-title', '01-Installation_de_l_environnement']],
  ['03-let_s_go', ['01-Les_bases', '02-if', '03-les_boucles', '04-Defer_Panic_Recover']],
  [
    '04-Let_s_Go_further',
    ['00-title', '01-Déclaration_de_types_et_structures', '02-Les_pointeurs', '03-Les_tableaux_slices_et_map'],
  ],
  ['05-Orienté_objet', ['00-title', '01-Les_méthodes', '02-Les_interfaces']],
  [
    '06-Programmation-concurrente',
    [
      '00-title',
      '01-Goroutine',
      '02-Channels',
      '03-Buffered-channels',
      '04-Channels-fermeture',
      '05-range',
      '06-exercice',
      '07-select',
      '08-exercice',
    ],
  ],
  ['07-et-maintenant', ['00-title', '01-aller-plus-loin', '02-merci']],
];

const go200 = [
  ['00-school', ['00-title', '01-speakers', '02-deroulement']],
  [
    '01-environnement',
    [
      '00-title',
      '01-installation',
      '02-workspace',
      '03-commandes',
      '04-makefile',
      '05-langage',
      '06-langage',
      '07-objectifs',
      '08-objectifs',
    ],
  ],
  [
    '02-CLI-args',
    [
      '00-title',
      '01-app-parameters',
      '02-stdlib-flag',
      '03-lib-cliv2',
      '04-lib-cliv2-code-flags',
      '05-lib-cliv2-code-main',
      '06-run',
      '07-exercise',
      '08-dependencies',
      '09-gopath',
      '10-vendor',
      '11-modules',
      '12-workflow',
    ],
  ],
  ['03-logging', ['00-title', '01-disclaimer', '02-stdlib-log', '03-lib-logrus', '04-logrus-code', '05-exercise']],
  [
    '04-concurrency',
    [
      '00-title',
      '01-explanations',
      '02-illustration',
      '03-struct',
      '04-constructor',
      '05-main-go-routine',
      '06-exercise',
    ],
  ],
];

const makeSlide = (dir) => (file) => ({ path: `${dir}/${file}.md` });
const pathReducerFactory = (baseDir = '') => (acc, [dir, files]) => [
  ...acc,
  ...files.map(makeSlide(`${baseDir}${dir}`)),
];

export const slides = go200
  .reduce(pathReducerFactory('go-200/'), [])
  .concat(go100.reduce(pathReducerFactory('go-100/'), []));
