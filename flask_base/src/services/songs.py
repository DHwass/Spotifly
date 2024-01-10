import json
import uuid
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user
from src.schemas.rating import RatingSchema

from src.schemas.song import SongSchema
from src.models.http_exceptions import *



songs_url = "http://localhost:8181/songs/"  # URL de l'API users (golang)

def get_song(id):
    response = requests.request(method="GET", url=songs_url+id)
    return response.json(), response.status_code

def get_songs():
    response = requests.request(method="GET", url=songs_url)
    return response.json(), response.status_code

def create_song(song_register):
    # on récupère le schéma song pour la requête vers l'API users
    song_schema = SongSchema().loads(json.dumps(song_register), unknown=EXCLUDE)
    print("this is from service ",song_schema)
    # on crée song côté API users
    response = requests.request(method="POST", url=songs_url, json=song_schema)

    if response.status_code != 200:
        return response.json(), response.status_code

  

    return response.json(), response.status_code

def update_song(id, song_update):


    # s'il y a quelque chose à changer côté API 
    song_schema = SongSchema().loads(json.dumps(song_update), unknown=EXCLUDE)
    print(song_schema)
    response = None
    if not SongSchema.is_empty(song_schema):
        # on lance la requête de modification
        response = requests.request(method="PUT", url=songs_url+id, json=song_schema)
        print(response.status_code)
        if response.status_code != 200:
            return response.json(), response.status_code

    return response.json(), response.status_code


def delete_song(id):
    response = requests.request(method="DELETE", url=songs_url+id)
    return  response.status_code
    
def get_ratings(id):
    response = requests.request(method="GET", url="https://ratings-alpha.edu.forestier.re/songs/"+id+"/ratings")
    return response.json(), response.status_code

def create_rating(id,rating_register):
  
    # on récupère le schéma rating pour la requête vers l'API ratings
    rating_schema = RatingSchema().loads(json.dumps(rating_register), unknown=EXCLUDE)
    print("here",rating_schema)
    #generation d'un uuid
    new_uuid = str(uuid.uuid4())

    rating_schema["id"]=new_uuid
    rating_schema["user_id"] = current_user.id
    rating_schema["rating_date"] = "2023/12/12"  
    rating_schema["song_id"] = id
    
    print(rating_schema)
    #check if music_id exists
    songs, status_code = get_song(id)
  

    if status_code != 200:
        return songs, status_code
    
    # on crée le rating côté API ratings
    response = requests.request(method="POST", url="https://ratings-alpha.edu.forestier.re/songs/"+id+"/ratings", json=rating_schema)
     
    print(response.json())

    if response.status_code != 201:
        return response.json(), response.status_code

  

    return response.json(), response.status_code

def get_rating(id,rating_id):
    response = requests.request(method="GET", url="https://ratings-alpha.edu.forestier.re/songs/"+id+"/ratings/"+rating_id)
    return response.json(), response.status_code

def update_rating(id,rating_id,rating_update):
    # s'il y a quelque chose à changer côté API 
    rating_schema = RatingSchema().loads(json.dumps(rating_update), unknown=EXCLUDE)
    print(rating_schema)
    response = None
    if not RatingSchema.is_empty(rating_schema):
        # on lance la requête de modification
        response = requests.request(method="PUT", url="https://ratings-alpha.edu.forestier.re/songs/"+id+"/ratings/"+rating_id, json=rating_schema)
        print(response.status_code)
        if response.status_code != 200:
            return response.json(), response.status_code

    return response.json(), response.status_code

def delete_rating(id,rating_id):
    response = requests.request(method="DELETE", url="https://ratings-alpha.edu.forestier.re/songs/"+id+"/ratings/"+rating_id)
    if response.status_code == 204:
        return "Deleted succesfully",204
    return  response.status_code,response.json()
