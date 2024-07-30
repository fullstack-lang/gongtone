import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongtonespecificComponent } from './gongtonespecific.component';

describe('GongtonespecificComponent', () => {
  let component: GongtonespecificComponent;
  let fixture: ComponentFixture<GongtonespecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongtonespecificComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(GongtonespecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
