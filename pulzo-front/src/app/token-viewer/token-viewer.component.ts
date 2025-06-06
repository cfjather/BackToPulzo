import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-token-viewer',
  imports: [CommonModule],
  standalone: true,
  templateUrl: './token-viewer.component.html',
  styleUrls: ['./token-viewer.component.css']
})
export class TokenViewerComponent implements OnInit {

  tokens: { token: string, usesLeft: number }[] = [];
  error: string | null = null;

  constructor(private http: HttpClient) {}

  ngOnInit() {
    this.loadTokens();
  }

  loadTokens() {
    this.error = null;
    this.tokens = [];

    this.http.get<{ [key: string]: { Uses: number } }>('http://localhost:8080/admin').subscribe({
        next: (data) => {
          this.tokens = Object.entries(data).map(([token, info]) => ({
            token,
            usesLeft: 5 - info.Uses
          }));
        },
        error: (err) => {
          console.error('Error cargando tokens:', err);
          this.error = 'Error cargando tokens activos.';
        }
      });
  }
}
