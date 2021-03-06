// source: alameda_api/v1alpha1/datahub/schemas/types.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var alameda_api_v1alpha1_datahub_common_metrics_pb = require('../../../../alameda_api/v1alpha1/datahub/common/metrics_pb.js');
goog.object.extend(proto, alameda_api_v1alpha1_datahub_common_metrics_pb);
var alameda_api_v1alpha1_datahub_common_types_pb = require('../../../../alameda_api/v1alpha1/datahub/common/types_pb.js');
goog.object.extend(proto, alameda_api_v1alpha1_datahub_common_types_pb);
goog.exportSymbol('proto.containersai.alameda.v1alpha1.datahub.schemas.Column', null, global);
goog.exportSymbol('proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement', null, global);
goog.exportSymbol('proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta', null, global);
goog.exportSymbol('proto.containersai.alameda.v1alpha1.datahub.schemas.Scope', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta.displayName = 'proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement.repeatedFields_, null);
};
goog.inherits(proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement.displayName = 'proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Column = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.containersai.alameda.v1alpha1.datahub.schemas.Column, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.containersai.alameda.v1alpha1.datahub.schemas.Column.displayName = 'proto.containersai.alameda.v1alpha1.datahub.schemas.Column';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta.prototype.toObject = function(opt_includeInstance) {
  return proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta.toObject = function(includeInstance, msg) {
  var f, obj = {
    scope: jspb.Message.getFieldWithDefault(msg, 1, 0),
    category: jspb.Message.getFieldWithDefault(msg, 2, ""),
    type: jspb.Message.getFieldWithDefault(msg, 3, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta;
  return proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.containersai.alameda.v1alpha1.datahub.schemas.Scope} */ (reader.readEnum());
      msg.setScope(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setCategory(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setType(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getScope();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getCategory();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getType();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
};


/**
 * optional Scope scope = 1;
 * @return {!proto.containersai.alameda.v1alpha1.datahub.schemas.Scope}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta.prototype.getScope = function() {
  return /** @type {!proto.containersai.alameda.v1alpha1.datahub.schemas.Scope} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.schemas.Scope} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta.prototype.setScope = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional string category = 2;
 * @return {string}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta.prototype.getCategory = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta.prototype.setCategory = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string type = 3;
 * @return {string}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta.prototype.getType = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta.prototype.setType = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement.repeatedFields_ = [5];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement.prototype.toObject = function(opt_includeInstance) {
  return proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, ""),
    metricType: jspb.Message.getFieldWithDefault(msg, 2, 0),
    resourceBoundary: jspb.Message.getFieldWithDefault(msg, 3, 0),
    resourceQuota: jspb.Message.getFieldWithDefault(msg, 4, 0),
    columnsList: jspb.Message.toObjectList(msg.getColumnsList(),
    proto.containersai.alameda.v1alpha1.datahub.schemas.Column.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement;
  return proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 2:
      var value = /** @type {!proto.containersai.alameda.v1alpha1.datahub.common.MetricType} */ (reader.readEnum());
      msg.setMetricType(value);
      break;
    case 3:
      var value = /** @type {!proto.containersai.alameda.v1alpha1.datahub.common.ResourceBoundary} */ (reader.readEnum());
      msg.setResourceBoundary(value);
      break;
    case 4:
      var value = /** @type {!proto.containersai.alameda.v1alpha1.datahub.common.ResourceQuota} */ (reader.readEnum());
      msg.setResourceQuota(value);
      break;
    case 5:
      var value = new proto.containersai.alameda.v1alpha1.datahub.schemas.Column;
      reader.readMessage(value,proto.containersai.alameda.v1alpha1.datahub.schemas.Column.deserializeBinaryFromReader);
      msg.addColumns(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getMetricType();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
  f = message.getResourceBoundary();
  if (f !== 0.0) {
    writer.writeEnum(
      3,
      f
    );
  }
  f = message.getResourceQuota();
  if (f !== 0.0) {
    writer.writeEnum(
      4,
      f
    );
  }
  f = message.getColumnsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      5,
      f,
      proto.containersai.alameda.v1alpha1.datahub.schemas.Column.serializeBinaryToWriter
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional containersai.alameda.v1alpha1.datahub.common.MetricType metric_type = 2;
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.MetricType}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement.prototype.getMetricType = function() {
  return /** @type {!proto.containersai.alameda.v1alpha1.datahub.common.MetricType} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.common.MetricType} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement.prototype.setMetricType = function(value) {
  return jspb.Message.setProto3EnumField(this, 2, value);
};


/**
 * optional containersai.alameda.v1alpha1.datahub.common.ResourceBoundary resource_boundary = 3;
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.ResourceBoundary}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement.prototype.getResourceBoundary = function() {
  return /** @type {!proto.containersai.alameda.v1alpha1.datahub.common.ResourceBoundary} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.common.ResourceBoundary} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement.prototype.setResourceBoundary = function(value) {
  return jspb.Message.setProto3EnumField(this, 3, value);
};


/**
 * optional containersai.alameda.v1alpha1.datahub.common.ResourceQuota resource_quota = 4;
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.ResourceQuota}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement.prototype.getResourceQuota = function() {
  return /** @type {!proto.containersai.alameda.v1alpha1.datahub.common.ResourceQuota} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.common.ResourceQuota} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement.prototype.setResourceQuota = function(value) {
  return jspb.Message.setProto3EnumField(this, 4, value);
};


/**
 * repeated Column columns = 5;
 * @return {!Array<!proto.containersai.alameda.v1alpha1.datahub.schemas.Column>}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement.prototype.getColumnsList = function() {
  return /** @type{!Array<!proto.containersai.alameda.v1alpha1.datahub.schemas.Column>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.containersai.alameda.v1alpha1.datahub.schemas.Column, 5));
};


/**
 * @param {!Array<!proto.containersai.alameda.v1alpha1.datahub.schemas.Column>} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement} returns this
*/
proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement.prototype.setColumnsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 5, value);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.schemas.Column=} opt_value
 * @param {number=} opt_index
 * @return {!proto.containersai.alameda.v1alpha1.datahub.schemas.Column}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement.prototype.addColumns = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 5, opt_value, proto.containersai.alameda.v1alpha1.datahub.schemas.Column, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Measurement.prototype.clearColumnsList = function() {
  return this.setColumnsList([]);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Column.prototype.toObject = function(opt_includeInstance) {
  return proto.containersai.alameda.v1alpha1.datahub.schemas.Column.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.containersai.alameda.v1alpha1.datahub.schemas.Column} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Column.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, ""),
    required: jspb.Message.getBooleanFieldWithDefault(msg, 2, false),
    columnType: jspb.Message.getFieldWithDefault(msg, 3, 0),
    dataType: jspb.Message.getFieldWithDefault(msg, 4, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.schemas.Column}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Column.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.containersai.alameda.v1alpha1.datahub.schemas.Column;
  return proto.containersai.alameda.v1alpha1.datahub.schemas.Column.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.containersai.alameda.v1alpha1.datahub.schemas.Column} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.schemas.Column}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Column.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 2:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setRequired(value);
      break;
    case 3:
      var value = /** @type {!proto.containersai.alameda.v1alpha1.datahub.common.ColumnType} */ (reader.readEnum());
      msg.setColumnType(value);
      break;
    case 4:
      var value = /** @type {!proto.containersai.alameda.v1alpha1.datahub.common.DataType} */ (reader.readEnum());
      msg.setDataType(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Column.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.containersai.alameda.v1alpha1.datahub.schemas.Column.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.containersai.alameda.v1alpha1.datahub.schemas.Column} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Column.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getRequired();
  if (f) {
    writer.writeBool(
      2,
      f
    );
  }
  f = message.getColumnType();
  if (f !== 0.0) {
    writer.writeEnum(
      3,
      f
    );
  }
  f = message.getDataType();
  if (f !== 0.0) {
    writer.writeEnum(
      4,
      f
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Column.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.schemas.Column} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Column.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional bool required = 2;
 * @return {boolean}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Column.prototype.getRequired = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 2, false));
};


/**
 * @param {boolean} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.schemas.Column} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Column.prototype.setRequired = function(value) {
  return jspb.Message.setProto3BooleanField(this, 2, value);
};


/**
 * optional containersai.alameda.v1alpha1.datahub.common.ColumnType column_type = 3;
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.ColumnType}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Column.prototype.getColumnType = function() {
  return /** @type {!proto.containersai.alameda.v1alpha1.datahub.common.ColumnType} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.common.ColumnType} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.schemas.Column} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Column.prototype.setColumnType = function(value) {
  return jspb.Message.setProto3EnumField(this, 3, value);
};


/**
 * optional containersai.alameda.v1alpha1.datahub.common.DataType data_type = 4;
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.DataType}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Column.prototype.getDataType = function() {
  return /** @type {!proto.containersai.alameda.v1alpha1.datahub.common.DataType} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.common.DataType} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.schemas.Column} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Column.prototype.setDataType = function(value) {
  return jspb.Message.setProto3EnumField(this, 4, value);
};


/**
 * @enum {number}
 */
proto.containersai.alameda.v1alpha1.datahub.schemas.Scope = {
  SCOPE_UNDEFINED: 0,
  SCOPE_APPLICATION: 1,
  SCOPE_EXECUTION: 2,
  SCOPE_METRIC: 3,
  SCOPE_PLANNING: 4,
  SCOPE_PREDICTION: 5,
  SCOPE_RECOMMENDATION: 6,
  SCOPE_RESOURCE: 7
};

goog.object.extend(exports, proto.containersai.alameda.v1alpha1.datahub.schemas);
