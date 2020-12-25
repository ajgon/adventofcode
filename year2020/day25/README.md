# Day 25: Combo Breaker

You finally reach the check-in desk. Unfortunately, their registration systems are currently offline, and they cannot check you in. Noticing the look on your face, they quickly add that tech support is already on the way! They even created all the room keys this morning; you can take yours now and give them your room deposit once the registration system comes back online.

The room key is a small [RFID](https://en.wikipedia.org/wiki/Radio-frequency_identification) card. Your room is on the 25th floor and the elevators are also temporarily out of service, so it takes what little energy you have left to even climb the stairs and navigate the halls. You finally reach the door to your room, swipe your card, and - __beep__ - the light turns red.

Examining the card more closely, you discover a phone number for tech support.

"Hello! How can we help you today?" You explain the situation.

"Well, it sounds like the card isn't sending the right command to unlock the door. If you go back to the check-in desk, surely someone there can reset it for you." Still catching your breath, you describe the status of the elevator and the exact number of stairs you just had to climb.

"I see! Well, your only other option would be to reverse-engineer the cryptographic handshake the card does with the door and then inject your own commands into the data stream, but that's definitely impossible." You thank them for their time.

Unfortunately for the door, you know a thing or two about cryptographic handshakes.

The handshake used by the card and the door involves an operation that __transforms__ a __subject number__. To transform a subject number, start with the value 1. Then, a number of times called the __loop size__, perform the following steps:

    Set the value to itself multiplied by the __subject number.__
    Set the value to the remainder after dividing the value by __`20201227`__.

The card always uses a specific, secret __loop size__ when it transforms a subject number. The door always uses a different, secret loop size.

The cryptographic handshake works like this:

- The __card__ transforms the subject number of __`7`__ according to the __card's__ secret loop size. The result is called the __card's public key.__
- The __door__ transforms the subject number of __`7`__ according to the __door's__ secret loop size. The result is called the __door's public key.__
- The card and door use the wireless RFID signal to transmit the two public keys (your puzzle input) to the other device. Now, the __card__ has the __door's__ public key, and the __door__ has the __card's__ public key. Because you can eavesdrop on the signal, you have both public keys, but neither device's loop size.
- The __card__ transforms the subject number of __the door's public key__ according to the __card's__ loop size. The result is the __encryption key.__
- The __door__ transforms the subject number of __the card's public key__ according to the __door's__ loop size. The result is the same __encryption key__ as the __card__ calculated.

If you can use the two public keys to determine each device's loop size, you will have enough information to calculate the secret __encryption key__ that the card and door use to communicate; this would let you send the `unlock` command directly to the door!

For example, suppose you know that the card's public key is `5764801`. With a little trial and error, you can work out that the card's loop size must be __`8`__, because transforming the initial subject number of `7` with a loop size of `8` produces `5764801`.

Then, suppose you know that the door's public key is `17807724`. By the same process, you can determine that the door's loop size is __`11`__, because transforming the initial subject number of `7` with a loop size of `11` produces `17807724`.

At this point, you can use either device's loop size with the other device's public key to calculate the __encryption key__. Transforming the subject number of `17807724` (the door's public key) with a loop size of `8` (the card's loop size) produces the encryption key, __`14897079`__. (Transforming the subject number of `5764801` (the card's public key) with a loop size of `11` (the door's loop size) produces the same encryption key: __`14897079`__.)

__What encryption key is the handshake trying to establish?__
