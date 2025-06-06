import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css'
})

export class LoginComponent {
  token: string | null = null;
  copied: boolean = false;

  constructor(private http: HttpClient) {}

  // Solicita un Token a GO
  getToken() {
    this.http.post<{ token: string }>('http://localhost:8080/login', {})
      .subscribe({
        next: (res) => {
          this.token = res.token;
          this.copied = false; // Limpiar Copia al obtener nuevo token
        },
        error: (err) => {
          console.error(err);
          alert('Error al obtener token');
        }
      });
  }

  // Copia el Token al Portapapeles
  copyToken() {
    if (!this.token) return;

    navigator.clipboard.writeText(this.token).then(() => {
      this.copied = true;
      setTimeout(() => {
        this.copied = false;
      }, 2000);
    }).catch(() => {
      alert('No se pudo copiar el token. Por favor cópialo manualmente.');
    });
  }
}