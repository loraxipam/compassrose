# compassrose
The compassrose library lets you convert compass angles to cardinal point directions up to 32 points. It provides descriptions for the cardinal directions in both English and Ancient Mediterranean in case you need to sail with Balboa or Columbus.

## Structures
*CompassRose* provides a ```[32][3]string``` of the compass directions and their descriptions  
*TwoPoints, FourPoints, EightPoints*... provide enum indexes to simplify using the "level" parameter in the below functions

## Functions
```DegreeToHeading(inDegrees float32, level int, standard bool) (shorthand string, text string)``` is the main function which expects an angle and a compass point "level" granularity parameter. It returns the direction "shorthand" such as "N" or "WSW" and the description of that shorthand. If you choose the "standard" description it will return boring English. If you set "standard" to false, it will return the swashbuckling goodness of pure maritime jargon known and loved by your favorite childhood pirates.

```DegreeToHeadingSouthfacing(inDegrees float32, level int, standard bool) (string, string)``` is your preferred function if you want to look at the world from a southern hemisphere state of mind. Then you can get these same compass point values using this helper function which will return the usual compass directions, but using a south-facing frame of reference. Using this function an ```inDegrees``` of 0.0 returns "S", not "N" as the above function provides.

## Example

Basic usage turns a floating point degree heading into descriptive text.  

A helper function for south-facing frames of reference returns absolute compass directions but with South being the zero direction. The reference frame still increments clockwise, though.

Both functions handle positive and negative degrees, as well as |values| greater than 360.

```golang
package main

import (
    "fmt"

    "github.com/loraxipam/compassrose"
)

func main(){
    var heading float32 = 237.49
    direction, jargon := compassrose.DegreeToHeading(heading, compassrose.SixteenPoints, false)
    fmt.Printf("On the poop deck of the Host, with Ruff Laarssen and Hump Van Waydown\n")
    fmt.Printf("heading %-.2f°, we hear this conversation:\n\n", heading)
    fmt.Printf("\"Argggh, Hump, keep the jibs full, heading %s.\"\n", jargon)
    fmt.Printf("\"Aye, sir, binnacle shows %s.\"\n\n", direction)

    heading = -456.0
    direction, jargon = compassrose.DegreeToHeading(heading, compassrose.ThirtyTwoPoints, true)
    fmt.Printf("Hump knows that a heading of -456.0° is %s, %s.\n",jargon, direction)
}
```

Results:

```
On the poop deck of the Host, with Ruff Laarssen and Hump Van Waydown
heading 237.49°, we hear this conversation:

"Argggh, Hump, keep the jibs full, heading Ponente Libeccio."
"Aye, sir, binnacle shows WSW."

Hump knows that a heading of -456.0° is West by South, WbS.


```
