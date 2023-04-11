package input

type RangeConverter struct {
	conversionMap map[Range]converter
}

type RangeType float64

func (r *RangeConverter) Convert(rangeId Range, invalue float64) Range {
	result, ok := r.conversionMap[rangeId]
	if ok {
		return Range(invalue)
	}
	return Range(result.convert(invalue))
}

type converter struct {
	MinInput  float64
	MaxInput  float64
	MinOutput float64
	MaxOutput float64
}

func (c *converter) convert(value float64) float64 {
	if value < c.MinInput {
		value = c.MinInput
	} else if value > c.MaxInput {
		value = c.MaxInput
	}

	interpolationFactor := (value - c.MinInput) / (c.MaxInput - c.MinInput)
	return (interpolationFactor * (c.MaxOutput - c.MinInput)) + c.MinOutput
}

type ConversionMap map[Range]converter

//class RangeConverter
//{
//// Internal helpers
//private:
//struct Converter
//{
//double MinimumInput;
//double MaximumInput;
//
//double MinimumOutput;
//double MaximumOutput;
//
//template <typename RangeType>
//RangeType Convert(RangeType invalue) const
//{
//double v = static_cast<double>(invalue);
//if(v < MinimumInput)
//v = MinimumInput;
//else if(v > MaximumInput)
//v = MaximumInput;
//
//double interpolationfactor = (v - MinimumInput) / (MaximumInput - MinimumInput);
//return static_cast<RangeType>((interpolationfactor * (MaximumOutput - MinimumOutput)) + MinimumOutput);
//}
//};
//
//// Internal type shortcuts
//private:
//typedef std::map<Range, Converter> ConversionMapT;
//
//// Construction
//public:
//explicit RangeConverter(std::wifstream& infile);
//
//// Conversion interface
//public:
//template <typename RangeType>
//RangeType Convert(Range rangeid, RangeType invalue) const
//{
//ConversionMapT::const_iterator iter = ConversionMap.find(rangeid);
//if(iter == ConversionMap.end())
//return invalue;
//
//return iter->second.Convert<RangeType>(invalue);
//}
//
//// Internal tracking
//private:
//ConversionMapT ConversionMap;
