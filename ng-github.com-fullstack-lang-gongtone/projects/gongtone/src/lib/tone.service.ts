// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs'
import { Observable, of } from 'rxjs'
import { catchError, map, tap } from 'rxjs/operators'

import { ToneAPI } from './tone-api'
import { Tone, CopyToneToToneAPI } from './tone'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class ToneService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  ToneServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private tonesUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.tonesUrl = origin + '/api/github.com/fullstack-lang/gongtone/go/v1/tones';
  }

  /** GET tones from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<ToneAPI[]> {
    return this.getTones(GONG__StackPath, frontRepo)
  }
  getTones(GONG__StackPath: string, frontRepo: FrontRepo): Observable<ToneAPI[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<ToneAPI[]>(this.tonesUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<ToneAPI[]>('getTones', []))
      );
  }

  /** GET tone by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<ToneAPI> {
    return this.getTone(id, GONG__StackPath, frontRepo)
  }
  getTone(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<ToneAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.tonesUrl}/${id}`;
    return this.http.get<ToneAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched tone id=${id}`)),
      catchError(this.handleError<ToneAPI>(`getTone id=${id}`))
    );
  }

  // postFront copy tone to a version with encoded pointers and post to the back
  postFront(tone: Tone, GONG__StackPath: string): Observable<ToneAPI> {
    let toneAPI = new ToneAPI
    CopyToneToToneAPI(tone, toneAPI)
    const id = typeof toneAPI === 'number' ? toneAPI : toneAPI.ID
    const url = `${this.tonesUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<ToneAPI>(url, toneAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<ToneAPI>('postTone'))
    );
  }
  
  /** POST: add a new tone to the server */
  post(tonedb: ToneAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<ToneAPI> {
    return this.postTone(tonedb, GONG__StackPath, frontRepo)
  }
  postTone(tonedb: ToneAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<ToneAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<ToneAPI>(this.tonesUrl, tonedb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted tonedb id=${tonedb.ID}`)
      }),
      catchError(this.handleError<ToneAPI>('postTone'))
    );
  }

  /** DELETE: delete the tonedb from the server */
  delete(tonedb: ToneAPI | number, GONG__StackPath: string): Observable<ToneAPI> {
    return this.deleteTone(tonedb, GONG__StackPath)
  }
  deleteTone(tonedb: ToneAPI | number, GONG__StackPath: string): Observable<ToneAPI> {
    const id = typeof tonedb === 'number' ? tonedb : tonedb.ID;
    const url = `${this.tonesUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<ToneAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted tonedb id=${id}`)),
      catchError(this.handleError<ToneAPI>('deleteTone'))
    );
  }

  // updateFront copy tone to a version with encoded pointers and update to the back
  updateFront(tone: Tone, GONG__StackPath: string): Observable<ToneAPI> {
    let toneAPI = new ToneAPI
    CopyToneToToneAPI(tone, toneAPI)
    const id = typeof toneAPI === 'number' ? toneAPI : toneAPI.ID
    const url = `${this.tonesUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<ToneAPI>(url, toneAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<ToneAPI>('updateTone'))
    );
  }

  /** PUT: update the tonedb on the server */
  update(tonedb: ToneAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<ToneAPI> {
    return this.updateTone(tonedb, GONG__StackPath, frontRepo)
  }
  updateTone(tonedb: ToneAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<ToneAPI> {
    const id = typeof tonedb === 'number' ? tonedb : tonedb.ID;
    const url = `${this.tonesUrl}/${id}`;


    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<ToneAPI>(url, tonedb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated tonedb id=${tonedb.ID}`)
      }),
      catchError(this.handleError<ToneAPI>('updateTone'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in ToneService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("ToneService" + error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {
    console.log(message)
  }
}