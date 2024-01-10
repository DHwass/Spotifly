import datetime
import json
import uuid
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user
import src.services.songs as songs_service

from src.schemas.rating import RatingSchema
from src.models.http_exceptions import *

ratings_url = "https://ratings-alpha.edu.forestier.re/songs/"  # URL de l'API ratings 

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
    songs, status_code = songs_service.get_song(id)
  

    if status_code != 200:
        return songs, status_code
    
    # on crée le rating côté API ratings
    response = requests.request(method="POST", url=ratings_url+id+"/ratings", json=rating_schema)
     
    print(response.json())

    if response.status_code != 201:
        return response.json(), response.status_code

  

    return response.json(), response.status_code

def get_ratings(id):
    response = requests.request(method="GET", url=ratings_url+id+"/ratings")
    return response.json(), response.status_code


