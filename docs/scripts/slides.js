// add speakers here and run : npm run prepare to make go-200 in sync with go-100
const speakers = [
  // comment speakers before presenting
  'SFR',
  'SDU',
  'OFU',
  'YFA',
  'YDA',
  'OGE',
  'APO',
  'FSA',
].map((trigram) => `speakers/${trigram}`);

// go100 & go200 : canonical structure for ordered 2-levels filesystem tree
// array of [directory, [file]] pairs
const go100 = [
  ['00-school', ['00-title', ...speakers]],
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
  ['00-school', ['00-title', ...speakers, '02-deroulement']],
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
  [
    '05-tests',
    [
      '00-presentation',
      '01-whitebox',
      '02-blackbox',
      // '03-test-counter' // invisible on source slide before conversion
      '04-coverage',
      '05-exercise',
    ],
  ],
  ['06-modelization', ['00-title', '01-struct', '02-enums', '03-struct-tags', '04-uuid', '05-dates', '06-exercise']],
  [
    '07-databases',
    [
      '00-title',
      '01-factory-explanations',
      '02-factory-illustration',
      '03-DAO-interface',
      '04-implementation-mock-01',
      '05-implementation-mock-02',
      '06-implementation-mongodb-01',
      '07-implementation-mongodb-02',
      '08-implementation-postgresql-01',
      '09-implementation-postgresql-02',
      '10-factory-01',
      '11-factory-02',
      '12-factory-03',
      '13-factory-04',
      '14-DAO-tests',
      '15-exercise',
    ],
  ],
  [
    '08-REST',
    [
      '00-title',
      '02-arch',
      '03-stdlib-http',
      '04-stdlib-http-handle',
      '05-stdlib-http-handler',
      '06-components',
      '07-compose-middlewares',
      '08-negroni',
      '09-statistics-middleware',
      '10-handlers-01',
      '11-handlers-02',
      '12-unit-tests',
      '13-end-to-end-tests',
      '14-subtests',
      '15-subtests-parallel',
      '16-exercise',
    ],
  ],
  [
    '09-tail',
    [
      '00-title',
      '01-benchmarks-01',
      '02-benchmarks-02',
      '03-profiling',
      '04-http-profiling',
      '05-docker',
      '06-poney',
      '07-well-done',
      '08-resources',
      '09-photo-credits',
      '10-links',
    ],
  ],
];

const makeSlidePath = (dir) => (file) => ({ path: `${dir}/${file}.md` });
const pathReducerFactory = (baseDir = '') => (acc, [dir, files]) => [
  ...acc,
  ...files.map(makeSlidePath(`${baseDir}${dir}`)),
];

const makeSlidePaths = ([slides, baseDir]) => slides.reduce(pathReducerFactory(baseDir), []);

// comment the presentation you don't want to use
export const slides = [
  [go100, 'go-100/'],
  [go200, 'go-200/'],
]
  .map(makeSlidePaths)
  .flat();
