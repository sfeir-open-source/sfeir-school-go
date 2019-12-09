# Les types de base
Les types de base de Go sont

    bool

    string

    int  int8  int16  int32  int64
    uint uint8 uint16 uint32 uint64 uintptr

    byte // alias pour uint8

    rune // alias pour int32. Représente un "code point" Unicode

    float32 float64

    complex64 complex128
    
L'exemple montre des variables de plusieurs types, et aussi que les déclarations de variables peuvent être « factorisé », en blocs, comme les imports.

Les types int, uint et uintptr sont généralement d'une largeur de 32 bits sur les systèmes 32 bits et 64 bits de large sur les systèmes 64 bits. Lorsque vous avez besoin d'une valeur entière vous devez utiliser int sauf si vous avez une raison particulière d'utiliser un type entier de taille définie ou non signé.

## Les valeurs zéro
Les variables déclarées sans valeur initiale explicite reçoivent leur valeur zéro.

La valeur zéro est :

- 0 pour les types numériques,
- False pour le type booléen, et
- "" (La chaîne vide) pour les chaînes.

## Les conversions de type
L'expression T(v) convertit la valeur v au type T.

Quelques conversions numériques :

    var i int = 42
    var f float64 = float64(i)
    var u uint = uint(f)
Ou, plus simplement :

    i := 42
    f := float64(i)
    u := uint(f)
Contrairement au C, en Go l'affectation entre les éléments de type différent exige une conversion explicite. Essayez de retirer le float64 ou uint des conversions dans l'exemple pour voir ce qui se passe.