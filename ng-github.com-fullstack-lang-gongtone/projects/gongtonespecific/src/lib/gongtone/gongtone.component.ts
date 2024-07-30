import { Component, OnInit } from '@angular/core';

import * as Tone from "tone"

import { MatIconModule } from '@angular/material/icon';
import { MatDividerModule } from '@angular/material/divider';
import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'lib-gongtone',
  standalone: true,
  imports: [
    MatButtonModule, MatDividerModule, MatIconModule
  ],
  templateUrl: './gongtone.component.html',
  styleUrl: './gongtone.component.css'
})
export class GongtoneComponent implements OnInit {

  synth: Tone.Synth<Tone.SynthOptions> | undefined

  ngOnInit(): void {
    this.synth = new Tone.Synth().toDestination()

  }


  onButtonClick() {

    if (this.synth) {
      //play a middle 'C' for the duration of an 8th note
      this.synth.triggerAttackRelease("C4", "8n")
      console.log("After trigger attack")

      Tone.start()
    }
  }

}
