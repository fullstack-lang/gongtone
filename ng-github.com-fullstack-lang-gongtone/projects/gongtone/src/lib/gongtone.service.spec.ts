import { TestBed } from '@angular/core/testing';

import { GongtoneService } from './gongtone.service';

describe('GongtoneService', () => {
  let service: GongtoneService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongtoneService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
