import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-character-viewer',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './character-viewer.component.html',
  styleUrls: ['./character-viewer.component.css']
})

export class CharacterViewerComponent implements OnInit {
  
  token: string = '';
  characters: any[] = [];
  usesLeft: number | null = null;
  error: string | null = null;

  constructor(private http: HttpClient) {}

  ngOnInit() {
    // No Characters Loaded without Token
  }

fetchCharacters() {
    this.error = null;
    this.characters = [];
    this.usesLeft = null;

    if (!this.token) {
      this.error = 'Por favor ingresa un Token.';
      return;
    }

    const headers = new HttpHeaders({
      'Authorization': this.token
    });

    this.http.get<any[]>('http://localhost:8080/characters', { headers })
    .subscribe({ 
      next: (data: any) => {
        this.usesLeft = data.usesLeft;
        this.characters = data.characters?.results || [];
      },
      error: (err) => {
        console.error('Request to API ERROR:', err);
        this.error = 'Lo sentimos, tu Token es Invalido o ha Vencido';
      }
    });
  }
}