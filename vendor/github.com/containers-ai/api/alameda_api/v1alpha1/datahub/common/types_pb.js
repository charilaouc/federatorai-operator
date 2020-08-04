// source: alameda_api/v1alpha1/datahub/common/types.proto
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

goog.exportSymbol('proto.containersai.alameda.v1alpha1.datahub.common.ColumnType', null, global);
goog.exportSymbol('proto.containersai.alameda.v1alpha1.datahub.common.DataType', null, global);
goog.exportSymbol('proto.containersai.alameda.v1alpha1.datahub.common.FunctionType', null, global);
goog.exportSymbol('proto.containersai.alameda.v1alpha1.datahub.common.ResourceBoundary', null, global);
goog.exportSymbol('proto.containersai.alameda.v1alpha1.datahub.common.ResourceQuota', null, global);
/**
 * @enum {number}
 */
proto.containersai.alameda.v1alpha1.datahub.common.ResourceBoundary = {
  RESOURCE_BOUNDARY_UNDEFINED: 0,
  RESOURCE_RAW: 1,
  RESOURCE_UPPER_BOUND: 2,
  RESOURCE_LOWER_BOUND: 3
};

/**
 * @enum {number}
 */
proto.containersai.alameda.v1alpha1.datahub.common.ResourceQuota = {
  RESOURCE_QUOTA_UNDEFINED: 0,
  RESOURCE_LIMIT: 1,
  RESOURCE_REQUEST: 2,
  RESOURCE_INITIAL_LIMIT: 3,
  RESOURCE_INITIAL_REQUEST: 4
};

/**
 * @enum {number}
 */
proto.containersai.alameda.v1alpha1.datahub.common.DataType = {
  DATATYPE_UNDEFINED: 0,
  DATATYPE_BOOL: 1,
  DATATYPE_INT: 2,
  DATATYPE_INT8: 3,
  DATATYPE_INT16: 4,
  DATATYPE_INT32: 5,
  DATATYPE_INT64: 6,
  DATATYPE_UINT: 7,
  DATATYPE_UINT8: 8,
  DATATYPE_UINT16: 9,
  DATATYPE_UINT32: 10,
  DATATYPE_UTIN64: 11,
  DATATYPE_FLOAT32: 12,
  DATATYPE_FLOAT64: 13,
  DATATYPE_STRING: 14
};

/**
 * @enum {number}
 */
proto.containersai.alameda.v1alpha1.datahub.common.ColumnType = {
  COLUMNTYPE_UDEFINED: 0,
  COLUMNTYPE_TAG: 1,
  COLUMNTYPE_FIELD: 2
};

/**
 * @enum {number}
 */
proto.containersai.alameda.v1alpha1.datahub.common.FunctionType = {
  FUNCTIONTYPE_UNDEFINED: 0,
  FUNCTIONTYPE_COUNT: 1,
  FUNCTIONTYPE_DISTINCT: 2,
  FUNCTIONTYPE_INTEGRAL: 3,
  FUNCTIONTYPE_MEAN: 4,
  FUNCTIONTYPE_MEDIAN: 5,
  FUNCTIONTYPE_MODE: 6,
  FUNCTIONTYPE_SPREAD: 7,
  FUNCTIONTYPE_STDDEV: 8,
  FUNCTIONTYPE_SUM: 9,
  FUNCTIONTYPE_BOTTOM: 10,
  FUNCTIONTYPE_FIRST: 11,
  FUNCTIONTYPE_LAST: 12,
  FUNCTIONTYPE_MAX: 13,
  FUNCTIONTYPE_MIN: 14,
  FUNCTIONTYPE_PERCENTILE: 15,
  FUNCTIONTYPE_SAMPLE: 16,
  FUNCTIONTYPE_TOP: 17,
  FUNCTIONTYPE_DERIVATIVE: 18
};

goog.object.extend(exports, proto.containersai.alameda.v1alpha1.datahub.common);
