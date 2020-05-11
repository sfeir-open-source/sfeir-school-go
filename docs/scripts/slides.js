// allFiles : canonical structure for ordered 2-levels filesystem tree
// array of [directory, [file]] pairs

const go100 = [
  ['00-school', ['00-TITLE', '01-speakers']],
  ['01-intro', ['00-TITLE', '01-Le_but_de_go']],
  ['02-installation', ['00-TITLE', '01-Installation_de_l_environnement']],
  ['03-let_s_go', ['01-Les_bases', '02-if', '03-les_boucles', '04-Defer_Panic_Recover']],
  [
    '04-Let_s_Go_further',
    ['00-TITLE', '01-Déclaration_de_types_et_structures', '02-Les_pointeurs', '03-Les_tableaux_slices_et_map'],
  ],
  ['05-Orienté_objet', ['00-TITLE', '01-Les_méthodes', '02-Les_interfaces']],
  [
    '06-Programmation-concurrente',
    [
      '00-TITLE',
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
  ['07-et-maintenant', ['00-TITLE', '01-aller-plus-loin', '02-merci']],
];

const go200 = [['00-school', ['00-TITLE', '01-speakers']]];

const makeSlide = (dir) => (file) => ({ path: `${dir}/${file}.md` });
const pathReducerFactory = (baseDir = '') => (acc, [dir, files]) => [
  ...acc,
  ...files.map(makeSlide(`${baseDir}${dir}`)),
];

export const slides = go100
  .reduce(pathReducerFactory('go-100/'), [])
  .concat(go200.reduce(pathReducerFactory('go-200/'), []));
