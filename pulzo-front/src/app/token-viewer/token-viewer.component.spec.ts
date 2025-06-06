import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TokenViewerComponent } from './token-viewer.component';

describe('TokenViewerComponent', () => {
  let component: TokenViewerComponent;
  let fixture: ComponentFixture<TokenViewerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TokenViewerComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TokenViewerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
