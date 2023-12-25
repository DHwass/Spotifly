from marshmallow import Schema, fields, validates_schema, ValidationError


# Schéma utilisateur de sortie (renvoyé au front)
class UserSchema(Schema):
    id = fields.String(description="UUID")
    #inscription_date = fields.DateTime(description="Inscription date")
    name = fields.String(description="Name")
    email = fields.String(description="email")
    
    @staticmethod
    def is_empty(obj):
        return (not obj.get("id") or obj.get("id") == "") and \
               (not obj.get("name") or obj.get("name") == "") and \
               (not obj.get("email") or obj.get("email") == "") 


class BaseUserSchema(Schema):
    name = fields.String(description="name")
    email = fields.String(description="email")
    password = fields.String(description="password")


# Schéma utilisateur de modification (name, email, password)
class UserUpdateSchema(BaseUserSchema):
    # permet de définir dans quelles conditions le schéma est validé ou nom
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("name" in data and data["name"] != "") or
                ("email" in data and data["email"] != "") or
                ("password" in data and data["password"] != "")):
            raise ValidationError("at least one of ['name','email','username','password'] must be specified")
