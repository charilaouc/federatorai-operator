// source: alameda_api/v1alpha1/datahub/data/types.proto
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
var alameda_api_v1alpha1_datahub_common_rawdata_pb = require('../../../../alameda_api/v1alpha1/datahub/common/rawdata_pb.js');
goog.object.extend(proto, alameda_api_v1alpha1_datahub_common_rawdata_pb);
var alameda_api_v1alpha1_datahub_common_types_pb = require('../../../../alameda_api/v1alpha1/datahub/common/types_pb.js');
goog.object.extend(proto, alameda_api_v1alpha1_datahub_common_types_pb);
var alameda_api_v1alpha1_datahub_schemas_types_pb = require('../../../../alameda_api/v1alpha1/datahub/schemas/types_pb.js');
goog.object.extend(proto, alameda_api_v1alpha1_datahub_schemas_types_pb);
goog.exportSymbol('proto.containersai.alameda.v1alpha1.datahub.data.Data', null, global);
goog.exportSymbol('proto.containersai.alameda.v1alpha1.datahub.data.Rawdata', null, global);
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
proto.containersai.alameda.v1alpha1.datahub.data.Data = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.containersai.alameda.v1alpha1.datahub.data.Data.repeatedFields_, null);
};
goog.inherits(proto.containersai.alameda.v1alpha1.datahub.data.Data, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.containersai.alameda.v1alpha1.datahub.data.Data.displayName = 'proto.containersai.alameda.v1alpha1.datahub.data.Data';
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
proto.containersai.alameda.v1alpha1.datahub.data.Rawdata = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.repeatedFields_, null);
};
goog.inherits(proto.containersai.alameda.v1alpha1.datahub.data.Rawdata, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.displayName = 'proto.containersai.alameda.v1alpha1.datahub.data.Rawdata';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.containersai.alameda.v1alpha1.datahub.data.Data.repeatedFields_ = [2];



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
proto.containersai.alameda.v1alpha1.datahub.data.Data.prototype.toObject = function(opt_includeInstance) {
  return proto.containersai.alameda.v1alpha1.datahub.data.Data.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.containersai.alameda.v1alpha1.datahub.data.Data} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.containersai.alameda.v1alpha1.datahub.data.Data.toObject = function(includeInstance, msg) {
  var f, obj = {
    schemaMeta: (f = msg.getSchemaMeta()) && alameda_api_v1alpha1_datahub_schemas_types_pb.SchemaMeta.toObject(includeInstance, f),
    rawdataList: jspb.Message.toObjectList(msg.getRawdataList(),
    proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.toObject, includeInstance)
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
 * @return {!proto.containersai.alameda.v1alpha1.datahub.data.Data}
 */
proto.containersai.alameda.v1alpha1.datahub.data.Data.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.containersai.alameda.v1alpha1.datahub.data.Data;
  return proto.containersai.alameda.v1alpha1.datahub.data.Data.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.containersai.alameda.v1alpha1.datahub.data.Data} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.data.Data}
 */
proto.containersai.alameda.v1alpha1.datahub.data.Data.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new alameda_api_v1alpha1_datahub_schemas_types_pb.SchemaMeta;
      reader.readMessage(value,alameda_api_v1alpha1_datahub_schemas_types_pb.SchemaMeta.deserializeBinaryFromReader);
      msg.setSchemaMeta(value);
      break;
    case 2:
      var value = new proto.containersai.alameda.v1alpha1.datahub.data.Rawdata;
      reader.readMessage(value,proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.deserializeBinaryFromReader);
      msg.addRawdata(value);
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
proto.containersai.alameda.v1alpha1.datahub.data.Data.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.containersai.alameda.v1alpha1.datahub.data.Data.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.containersai.alameda.v1alpha1.datahub.data.Data} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.containersai.alameda.v1alpha1.datahub.data.Data.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSchemaMeta();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      alameda_api_v1alpha1_datahub_schemas_types_pb.SchemaMeta.serializeBinaryToWriter
    );
  }
  f = message.getRawdataList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.serializeBinaryToWriter
    );
  }
};


/**
 * optional containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta schema_meta = 1;
 * @return {?proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta}
 */
proto.containersai.alameda.v1alpha1.datahub.data.Data.prototype.getSchemaMeta = function() {
  return /** @type{?proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta} */ (
    jspb.Message.getWrapperField(this, alameda_api_v1alpha1_datahub_schemas_types_pb.SchemaMeta, 1));
};


/**
 * @param {?proto.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta|undefined} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.data.Data} returns this
*/
proto.containersai.alameda.v1alpha1.datahub.data.Data.prototype.setSchemaMeta = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.data.Data} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.data.Data.prototype.clearSchemaMeta = function() {
  return this.setSchemaMeta(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.containersai.alameda.v1alpha1.datahub.data.Data.prototype.hasSchemaMeta = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * repeated Rawdata rawdata = 2;
 * @return {!Array<!proto.containersai.alameda.v1alpha1.datahub.data.Rawdata>}
 */
proto.containersai.alameda.v1alpha1.datahub.data.Data.prototype.getRawdataList = function() {
  return /** @type{!Array<!proto.containersai.alameda.v1alpha1.datahub.data.Rawdata>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.containersai.alameda.v1alpha1.datahub.data.Rawdata, 2));
};


/**
 * @param {!Array<!proto.containersai.alameda.v1alpha1.datahub.data.Rawdata>} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.data.Data} returns this
*/
proto.containersai.alameda.v1alpha1.datahub.data.Data.prototype.setRawdataList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.data.Rawdata=} opt_value
 * @param {number=} opt_index
 * @return {!proto.containersai.alameda.v1alpha1.datahub.data.Rawdata}
 */
proto.containersai.alameda.v1alpha1.datahub.data.Data.prototype.addRawdata = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.containersai.alameda.v1alpha1.datahub.data.Rawdata, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.data.Data} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.data.Data.prototype.clearRawdataList = function() {
  return this.setRawdataList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.repeatedFields_ = [5];



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
proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.prototype.toObject = function(opt_includeInstance) {
  return proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.containersai.alameda.v1alpha1.datahub.data.Rawdata} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.toObject = function(includeInstance, msg) {
  var f, obj = {
    measurement: jspb.Message.getFieldWithDefault(msg, 1, ""),
    metricType: jspb.Message.getFieldWithDefault(msg, 2, 0),
    resourceBoundary: jspb.Message.getFieldWithDefault(msg, 3, 0),
    resourceQuota: jspb.Message.getFieldWithDefault(msg, 4, 0),
    groupsList: jspb.Message.toObjectList(msg.getGroupsList(),
    alameda_api_v1alpha1_datahub_common_rawdata_pb.Group.toObject, includeInstance)
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
 * @return {!proto.containersai.alameda.v1alpha1.datahub.data.Rawdata}
 */
proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.containersai.alameda.v1alpha1.datahub.data.Rawdata;
  return proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.containersai.alameda.v1alpha1.datahub.data.Rawdata} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.data.Rawdata}
 */
proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setMeasurement(value);
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
      var value = new alameda_api_v1alpha1_datahub_common_rawdata_pb.Group;
      reader.readMessage(value,alameda_api_v1alpha1_datahub_common_rawdata_pb.Group.deserializeBinaryFromReader);
      msg.addGroups(value);
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
proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.containersai.alameda.v1alpha1.datahub.data.Rawdata} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMeasurement();
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
  f = message.getGroupsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      5,
      f,
      alameda_api_v1alpha1_datahub_common_rawdata_pb.Group.serializeBinaryToWriter
    );
  }
};


/**
 * optional string measurement = 1;
 * @return {string}
 */
proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.prototype.getMeasurement = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.data.Rawdata} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.prototype.setMeasurement = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional containersai.alameda.v1alpha1.datahub.common.MetricType metric_type = 2;
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.MetricType}
 */
proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.prototype.getMetricType = function() {
  return /** @type {!proto.containersai.alameda.v1alpha1.datahub.common.MetricType} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.common.MetricType} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.data.Rawdata} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.prototype.setMetricType = function(value) {
  return jspb.Message.setProto3EnumField(this, 2, value);
};


/**
 * optional containersai.alameda.v1alpha1.datahub.common.ResourceBoundary resource_boundary = 3;
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.ResourceBoundary}
 */
proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.prototype.getResourceBoundary = function() {
  return /** @type {!proto.containersai.alameda.v1alpha1.datahub.common.ResourceBoundary} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.common.ResourceBoundary} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.data.Rawdata} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.prototype.setResourceBoundary = function(value) {
  return jspb.Message.setProto3EnumField(this, 3, value);
};


/**
 * optional containersai.alameda.v1alpha1.datahub.common.ResourceQuota resource_quota = 4;
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.ResourceQuota}
 */
proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.prototype.getResourceQuota = function() {
  return /** @type {!proto.containersai.alameda.v1alpha1.datahub.common.ResourceQuota} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.common.ResourceQuota} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.data.Rawdata} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.prototype.setResourceQuota = function(value) {
  return jspb.Message.setProto3EnumField(this, 4, value);
};


/**
 * repeated containersai.alameda.v1alpha1.datahub.common.Group groups = 5;
 * @return {!Array<!proto.containersai.alameda.v1alpha1.datahub.common.Group>}
 */
proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.prototype.getGroupsList = function() {
  return /** @type{!Array<!proto.containersai.alameda.v1alpha1.datahub.common.Group>} */ (
    jspb.Message.getRepeatedWrapperField(this, alameda_api_v1alpha1_datahub_common_rawdata_pb.Group, 5));
};


/**
 * @param {!Array<!proto.containersai.alameda.v1alpha1.datahub.common.Group>} value
 * @return {!proto.containersai.alameda.v1alpha1.datahub.data.Rawdata} returns this
*/
proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.prototype.setGroupsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 5, value);
};


/**
 * @param {!proto.containersai.alameda.v1alpha1.datahub.common.Group=} opt_value
 * @param {number=} opt_index
 * @return {!proto.containersai.alameda.v1alpha1.datahub.common.Group}
 */
proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.prototype.addGroups = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 5, opt_value, proto.containersai.alameda.v1alpha1.datahub.common.Group, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.containersai.alameda.v1alpha1.datahub.data.Rawdata} returns this
 */
proto.containersai.alameda.v1alpha1.datahub.data.Rawdata.prototype.clearGroupsList = function() {
  return this.setGroupsList([]);
};


goog.object.extend(exports, proto.containersai.alameda.v1alpha1.datahub.data);
