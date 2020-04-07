
function schoolSlides() {
  return ['00-school/00-TITLE.md',
    '00-school/01-speakers.md'];
}

function introSlides() {
  return ['01-intro/00-TITLE.md',
    '01-intro/01-Le_but_de_go.md'];
}

function installationSlides() {
  return ['02-installation/00-TITLE.md',
    '02-installation/01-Installation_de_l_environnement.md'];
}

function basicsSlides() {
  return ['03-let_s_go/00-TITLE.md',
    '03-let_s_go/01-Les_bases.md',
    '03-let_s_go/02-if.md',
    '03-let_s_go/03-les_boucles.md',
    '03-let_s_go/04-Defer_Panic_Recover.md'];
}

function typesSlides() {
  return ['04-Let_s_Go_further/00-TITLE.md',
    '04-Let_s_Go_further/01-Déclaration_de_types_et_structures.md',
    '04-Let_s_Go_further/02-Les_pointeurs.md',
    '04-Let_s_Go_further/03-Les_tableaux_slices_et_map.md'];
}

function objOrientedSlides() {
  return ['05-Orienté_objet/00-TITLE.md',
    '05-Orienté_objet/01-Les_méthodes.md',
    '05-Orienté_objet/02-Les_interfaces.md'];
}

function concurrencySlides() {
  return ['06-Programmation-concurrente/00-TITLE.md',
    '06-Programmation-concurrente/01-Goroutine.md',
    '06-Programmation-concurrente/02-Channels.md',
    '06-Programmation-concurrente/03-Buffered-channels.md',
    '06-Programmation-concurrente/04-Channels-fermeture.md',
    '06-Programmation-concurrente/05-range.md',
    '06-Programmation-concurrente/06-exercice.md',
    '06-Programmation-concurrente/07-select.md',
    '06-Programmation-concurrente/08-exercice.md'
  ];
}

function finalSlides() {
  return ['07-et-maintenant/00-TITLE.md',
    '07-et-maintenant/01-aller-plus-loin.md',
    '07-et-maintenant/02-merci.md'
  ];
}

function formation() {
  return [
    ...schoolSlides(),
    ...introSlides(),
    ...installationSlides(),
    ...basicsSlides(),
    ...typesSlides(),
    ...objOrientedSlides(),
    ...concurrencySlides(),
    ...finalSlides()
  ].map(slidePath => {
    return { path: slidePath };
  });
}

export function usedSlides() {
  return formation();
}

