<!-- .slide: class="sfeir-bg-white-3" -->

# Avant de se lancer

<b>Le but de Go</b>
<ul>
    <li>Rassembler les bonnes idées provenant d’autres langages,</li>
    <li>Abolir les fonctionnalités débouchant sur un code complexe et peu fiable…</li>
</ul>
<p>...dans le but d’écrire du code simple, expressif, robuste et efficace.</p>

##--##

# Avant de se lancer

<b>Un langage généraliste</b>
<ul>
    <li>Comme le C, il sera adapté dans pratiquement tous les domaines de programmation.</li>
    <li>Idéal pour le cloud.</li>
    <li>Déjà utilisé dans le graphisme, les applications mobiles, le machine learning, WASM, ...</li>
</ul>

##--##

# Avant de se lancer

<b>Déjà adopté par des grands</b>
<ul>
    <li>Google</li>
    <li>Docker</li>
    <li>Kubernetes</li>
    <li>Dropbox</li>
    <li>Spotify</li>
    <li>Hashicorp</li>
    <li>SoundCloud</li>
    <li>etc...</li>
</ul>

##--##

# Le language

<b>Déjà adopté par des grands</b>
<ul>
    <li>Né en 2009 chez Google (après les processeurs multi-coeurs) et OSS</li>
    <li>Binaire compilé autoporteur (début plugin depuis Go 1.8)</li>
    <li>Orienté objet</li>
    <li>Garbage collector (sub millisecond pour 17 Go de heap)</li>
    <li>Pointeurs 😱</li>
    <li>Goroutines</li>
        <ul>
            <li>Assimilable à un thread</li>
            <li>Mais ce n’est <b>PAS</b> un thread ⇒ <b>beaucoup plus léger</b></li>
        </ul>
    <li>Channels</li>
    <ul>
        <li><b>Do not communicate by sharing memory; share memory by communicating.</b></li>
        <li>Synchronisation</li>
        <li>Multiplexage (<b>select</b>)</li>
    </ul>
</ul>

##--##

# Le language

<b>Les mots clés</b>
<ul>
    <li><b>Dépendances :</b> import package</li>
    <li><b>Conditionnelles :</b> if else switch case fallthrough break default goto select</li>
    <li><b>Itérations :</b> for range continue</li>
    <li><b>Type :</b> var func interface struct chan const type map make</li>
    <li><b>Misc :</b> defer go return panic recover</li>
</ul>
![h-350](./assets/images/mots_clés.JPG)<!-- .element: class="bottom-image" -->
![h-350](./assets/images/i_know.JPG)<!-- .element: class="bottom-image" -->






