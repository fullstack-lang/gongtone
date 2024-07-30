import { TestBed } from '@angular/core/testing';

import { GongtonespecificService } from './gongtonespecific.service';

describe('GongtonespecificService', () => {
  let service: GongtonespecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongtonespecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
