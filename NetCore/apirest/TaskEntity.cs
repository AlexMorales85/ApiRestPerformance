using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;

namespace apirest
{
    [BsonIgnoreExtraElements(true)]
    public class TaskEntity
    {
        [BsonId]
        [BsonRepresentation(BsonType.ObjectId)]
        public string Id { get; set; }
        [BsonElement("title")]
        public string Title { get; set; }
        [BsonElement("description")]
        public string Description { get; set; }
    }
}
