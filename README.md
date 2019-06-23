# algdatii-snake

## About

Unser Programm lässt Snake durch eine KI spielen, welche bei einer Entscheidung von genau 2 Möglichkeiten den weiteren Weg simuliert und anhand dessen entscheidet.

## Getting started

Das Repository kann mithilfe des folgenden Befehls geklont werden:

```
$ git clone https://github.com/ob-algdatii-ss19/leistungsnachweis-asiansneverdie.git
```

ins Verzeichnis:

```
$ cd leistungsnachweis-asiansneverdie
```

und ausführen mit:

```
$ go run main.go
```

## Algorithmus

Der Algorithmus hat 4 Fälle.
* 0 Wege in die die Schlange gehen kann resultiert im Ende. (Sie ist gestorben)
* 1 Weg -> Sie muss den Weg laufen
* 2 Wege -> Unsere Simulation kommt zum Einsatz. Wir simulieren, ob die Schlange ihre komplette Länge überlebt, wenn sie sich für 1 der beiden Wege entscheidet. Falls 
** ja, dann läuft sie diesen Weg
** nein, dann den anderen
* Falls sie in der Simulation wieder auf 2 Möglichkeiten trifft, entscheidet sie wieder wie oben. Das ganze geschieht bis zu einer Tiefe von 3.
* 3 Wege -> Die Schlange läuft den schnellsten Weg zum Essen.

## Usage

Nach ausführen der main wird nun in der Console das Snake Spiel angezeigt.

* Der Kreis simbolisiert den Kopf der Snake.
* Rauten sind der Spielfeldrand. 
* Das Plus Symbol ist der Körper der Schlange.
* $ ist das Symbol für Essen.

Eine Aufnahme aus dem Spiel könnte beispielsweise so aussehen:

![Screenshot](https://github.com/ob-algdatii-ss19/leistungsnachweis-asiansneverdie/tree/feature/readme/snake_game.png)
