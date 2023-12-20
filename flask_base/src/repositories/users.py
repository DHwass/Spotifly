from src.helpers import db
from src.models.user import User


def get_user(email):
    return db.session.query(User).filter(User.email == email).first()


def get_user_from_id(id):
    return User.query.get(id)


def add_user(user):
    db.session.add(user)
    db.session.commit()


def update_user(user):
    existing_user = get_user_from_id(user.id)
    existing_user.username = user.username
    existing_user.Name = user.name
    existing_user.encrypted_password = user.encrypted_password
    db.session.commit()


def delete_user(id):
    db.session.delete(get_user_from_id(id))
    db.session.commit()
